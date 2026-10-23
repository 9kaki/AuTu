package core

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/metacubex/http"
	"github.com/metacubex/mihomo/component/ca"
)

const (
	defaultSpeedTestTimeout = 8 * time.Second
	defaultDownloadTestURL  = "https://speed.cloudflare.com/__down?bytes=25000000"
	defaultUploadTestURL    = "https://speed.cloudflare.com/__up"
	downloadTestDuration    = 2500 * time.Millisecond
	uploadTestDuration      = 2000 * time.Millisecond
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
		TLSHandshakeTimeout: 3 * time.Second,
		MaxIdleConns:        1,
		IdleConnTimeout:     3 * time.Second,
	}
	defer transport.CloseIdleConnections()

	client := &http.Client{
		Transport: transport,
	}

	testDownload := params.Mode == "" || params.Mode == "both" || params.Mode == "download"
	testUpload := params.Mode == "both" || params.Mode == "upload"

	perPhaseDuration := downloadTestDuration
	if totalTimeout < 6*time.Second {
		perPhaseDuration = 1800 * time.Millisecond
	}

	if testDownload {
		runDownloadTest(client, downloadURL, result, perPhaseDuration)
	}

	if testUpload {
		runUploadTest(client, uploadURL, result, perPhaseDuration)
	}

	return result
}

func runDownloadTest(client *http.Client, url string, result *SpeedTestResult, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration+2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		if result.Error == "" {
			result.Error = err.Error()
		}
		return
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		if result.Error == "" {
			result.Error = err.Error()
		}
		return
	}
	defer resp.Body.Close()

	if result.Delay == 0 {
		result.Delay = time.Since(start).Milliseconds()
	}

	buf := make([]byte, 32*1024)
	var totalBytes int64
	var peakSpeed int64
	var windowBytes int64
	windowStart := time.Now()
	startTime := time.Now()

	for {
		if time.Since(startTime) >= duration {
			break
		}
		n, rErr := resp.Body.Read(buf)
		if n > 0 {
			totalBytes += int64(n)
			windowBytes += int64(n)
		}

		wElapsed := time.Since(windowStart)
		if wElapsed >= 300*time.Millisecond {
			sec := wElapsed.Seconds()
			if sec > 0 {
				curSpeed := int64(float64(windowBytes) / sec)
				if curSpeed > peakSpeed {
					peakSpeed = curSpeed
				}
			}
			windowBytes = 0
			windowStart = time.Now()
		}

		if rErr != nil {
			break
		}
	}

	elapsed := time.Since(startTime).Seconds()
	var avgSpeed int64
	if elapsed > 0 && totalBytes > 0 {
		avgSpeed = int64(float64(totalBytes) / elapsed)
	}

	if peakSpeed > avgSpeed {
		result.DownloadSpeed = peakSpeed
	} else {
		result.DownloadSpeed = avgSpeed
	}
}

func runUploadTest(client *http.Client, url string, result *SpeedTestResult, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration+2*time.Second)
	defer cancel()

	reader := &countedZeroReader{
		remain:   50 * 1024 * 1024,
		duration: duration,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reader)
	if err != nil {
		if result.Error == "" {
			result.Error = err.Error()
		}
		return
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := client.Do(req)
	if err != nil {
		// Even if Do returned error on cancel/EOF, reader.readBytes counts what was uploaded
	} else {
		defer resp.Body.Close()
		io.Copy(io.Discard, resp.Body)
	}

	elapsed := time.Since(reader.startTime).Seconds()
	if elapsed > 0 && reader.readBytes > 0 {
		result.UploadSpeed = int64(float64(reader.readBytes) / elapsed)
	}
}
