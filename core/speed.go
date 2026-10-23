package core

import (
	"context"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/metacubex/http"
	"github.com/metacubex/mihomo/component/ca"
)

const (
	defaultSpeedTestTimeout = 10 * time.Second
	defaultDownloadTestURL  = "https://dl.google.com/chrome/mac/universal/stable/GGRO/googlechrome.dmg"
	defaultUploadTestURL    = "https://speed.cloudflare.com/__up"
	downloadTestDuration    = 3500 * time.Millisecond
	uploadTestDuration      = 2500 * time.Millisecond
	downloadConcurrency     = 4
	uploadConcurrency       = 2
	downloadBufferSize      = 128 * 1024
)

type countedZeroReader struct {
	remain    int64
	readBytes int64
	startTime time.Time
	duration  time.Duration
}

func (r *countedZeroReader) Read(p []byte) (n int, err error) {
	if !r.startTime.IsZero() && time.Since(r.startTime) >= r.duration {
		return 0, io.EOF
	}
	if r.remain <= 0 {
		return 0, io.EOF
	}
	if r.startTime.IsZero() {
		r.startTime = time.Now()
	}
	toRead := int64(len(p))
	if toRead > r.remain {
		toRead = r.remain
	}
	for i := int64(0); i < toRead; i++ {
		p[i] = 0
	}
	r.remain -= toRead
	r.readBytes += toRead
	return int(toRead), nil
}

func handleSpeedTest(params *SpeedTestParams) *SpeedTestResult {
	result := &SpeedTestResult{
		ProxyName: params.ProxyName,
	}

	proxy := lookupProxy(params.ProxyName)
	if proxy == nil {
		result.Error = fmt.Sprintf("proxy %q not found", params.ProxyName)
		return result
	}

	downloadURL := params.DownloadURL
	if downloadURL == "" {
		downloadURL = defaultDownloadTestURL
	}
	uploadURL := params.UploadURL
	if uploadURL == "" {
		uploadURL = defaultUploadTestURL
	}

	totalTimeout := defaultSpeedTestTimeout
	if params.Timeout > 0 {
		totalTimeout = time.Duration(params.Timeout) * time.Millisecond
	}

	dialer := &probeDialer{proxy: proxy}
	tlsConfig, _ := ca.GetTLSConfig(ca.Option{})

	transport := &http.Transport{
		DialContext:         dialer.dial,
		TLSClientConfig:     tlsConfig,
		TLSHandshakeTimeout: 4 * time.Second,
		MaxIdleConns:        16,
		MaxIdleConnsPerHost: 8,
		IdleConnTimeout:     5 * time.Second,
	}
	defer transport.CloseIdleConnections()

	client := &http.Client{
		Transport: transport,
	}

	testDownload := params.Mode == "" || params.Mode == "both" || params.Mode == "download"
	testUpload := params.Mode == "both" || params.Mode == "upload"

	perPhaseDuration := downloadTestDuration
	if totalTimeout < 7*time.Second {
		perPhaseDuration = 2000 * time.Millisecond
	}

	if testDownload {
		runDownloadTest(client, downloadURL, result, perPhaseDuration)
		if result.DownloadSpeed <= 0 && downloadURL != defaultDownloadTestURL {
			result.Error = ""
			runDownloadTest(client, defaultDownloadTestURL, result, perPhaseDuration)
		}
	}

	if testUpload {
		runUploadTest(client, uploadURL, result, perPhaseDuration)
	}

	return result
}

func runDownloadTest(client *http.Client, url string, result *SpeedTestResult, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration+3*time.Second)
	defer cancel()

	var (
		totalBytes   int64
		windowBytes  int64
		peakSpeed    int64
		firstDelayMs int64
		errOnce      sync.Once
		wg           sync.WaitGroup
	)

	startTime := time.Now()

	ticker := time.NewTicker(200 * time.Millisecond)
	tickerDone := make(chan struct{})
	go func() {
		lastSample := time.Now()
		for {
			select {
			case <-tickerDone:
				ticker.Stop()
				return
			case t := <-ticker.C:
				dt := t.Sub(lastSample).Seconds()
				wb := atomic.SwapInt64(&windowBytes, 0)
				if dt > 0 && wb > 0 {
					cur := int64(float64(wb) / dt)
					for {
						old := atomic.LoadInt64(&peakSpeed)
						if cur <= old || atomic.CompareAndSwapInt64(&peakSpeed, old, cur) {
							break
						}
					}
				}
				lastSample = t
			}
		}
	}()

	worker := func() {
		defer wg.Done()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			errOnce.Do(func() {
				if result.Error == "" {
					result.Error = err.Error()
				}
			})
			return
		}
		req.Header.Set("User-Agent", browserUserAgent)
		req.Header.Set("Accept", "*/*")

		reqStart := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			errOnce.Do(func() {
				if result.Error == "" {
					result.Error = err.Error()
				}
			})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			errOnce.Do(func() {
				if result.Error == "" {
					result.Error = fmt.Sprintf("HTTP %d (%s)", resp.StatusCode, resp.Status)
				}
			})
			return
		}

		reqDelay := time.Since(reqStart).Milliseconds()
		if atomic.CompareAndSwapInt64(&firstDelayMs, 0, reqDelay) {
			if result.Delay == 0 {
				result.Delay = reqDelay
			}
		}

		buf := make([]byte, downloadBufferSize)
		for {
			if time.Since(startTime) >= duration {
				break
			}
			n, rErr := resp.Body.Read(buf)
			if n > 0 {
				atomic.AddInt64(&totalBytes, int64(n))
				atomic.AddInt64(&windowBytes, int64(n))
			}
			if rErr != nil {
				break
			}
		}
	}

	wg.Add(downloadConcurrency)
	for i := 0; i < downloadConcurrency; i++ {
		go worker()
	}

	wg.Wait()
	close(tickerDone)

	elapsed := time.Since(startTime).Seconds()
	tot := atomic.LoadInt64(&totalBytes)
	var avgSpeed int64
	if elapsed > 0 && tot > 0 {
		avgSpeed = int64(float64(tot) / elapsed)
	}

	peak := atomic.LoadInt64(&peakSpeed)
	if peak > avgSpeed {
		result.DownloadSpeed = peak
	} else {
		result.DownloadSpeed = avgSpeed
	}
}

func runUploadTest(client *http.Client, url string, result *SpeedTestResult, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration+3*time.Second)
	defer cancel()

	var (
		totalUploadedBytes int64
		firstStart         time.Time
		startMu            sync.Mutex
		wg                 sync.WaitGroup
	)

	worker := func() {
		defer wg.Done()
		reader := &countedZeroReader{
			remain:   50 * 1024 * 1024,
			duration: duration,
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reader)
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/octet-stream")
		req.Header.Set("User-Agent", browserUserAgent)

		startMu.Lock()
		if firstStart.IsZero() {
			firstStart = time.Now()
		}
		startMu.Unlock()

		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			io.Copy(io.Discard, resp.Body)
		}
		atomic.AddInt64(&totalUploadedBytes, reader.readBytes)
	}

	wg.Add(uploadConcurrency)
	for i := 0; i < uploadConcurrency; i++ {
		go worker()
	}

	wg.Wait()

	startMu.Lock()
	t0 := firstStart
	startMu.Unlock()

	if !t0.IsZero() {
		elapsed := time.Since(t0).Seconds()
		tot := atomic.LoadInt64(&totalUploadedBytes)
		if elapsed > 0 && tot > 0 {
			result.UploadSpeed = int64(float64(tot) / elapsed)
		}
	}
}
