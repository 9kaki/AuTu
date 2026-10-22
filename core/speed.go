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
	defaultSpeedTestTimeout = 10 * time.Second
	defaultDownloadTestURL  = "https://speed.cloudflare.com/__down?bytes=25000000"
	defaultUploadTestURL    = "https://speed.cloudflare.com/__up"
	speedTestDuration       = 4 * time.Second
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

	dialer := &probeDialer{proxy: proxy}
	tlsConfig, _ := ca.GetTLSConfig(ca.Option{})

	transport := &http.Transport{
		DialContext:         dialer.dial,
		TLSClientConfig:     tlsConfig,
		TLSHandshakeTimeout: 5 * time.Second,
		MaxIdleConns:        1,
		IdleConnTimeout:     5 * time.Second,
	}
	defer transport.CloseIdleConnections()

	client := &http.Client{
		Transport: transport,
	}

	testDownload := params.Mode == "" || params.Mode == "both" || params.Mode == "download"
	testUpload := params.Mode == "both" || params.Mode == "upload"

	if testDownload {
		runDownloadTest(client, downloadURL, result, speedTestDuration)
	}

	if testUpload {
		runUploadTest(client, uploadURL, result, speedTestDuration)
	}

	return result
}

func runDownloadTest(client *http.Client, url string, result *SpeedTestResult, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration+3*time.Second)
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
	startTime := time.Now()

	for {
		if time.Since(startTime) >= duration {
			break
		}
		n, rErr := resp.Body.Read(buf)
		if n > 0 {
			totalBytes += int64(n)
		}
		if rErr != nil {
			break
		}
	}

	elapsed := time.Since(startTime).Seconds()
	if elapsed > 0 && totalBytes > 0 {
		result.DownloadSpeed = int64(float64(totalBytes) / elapsed)
	}
}

func runUploadTest(client *http.Client, url string, result *SpeedTestResult, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration+3*time.Second)
	defer cancel()

	reader := &countedZeroReader{
		remain:   30 * 1024 * 1024,
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
		// counted bytes
	} else {
		defer resp.Body.Close()
		io.Copy(io.Discard, resp.Body)
	}

	elapsed := time.Since(reader.startTime).Seconds()
	if elapsed > 0 && reader.readBytes > 0 {
		result.UploadSpeed = int64(float64(reader.readBytes) / elapsed)
	}
}
