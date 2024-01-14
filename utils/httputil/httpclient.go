package httputil

import (
	"net/http"
	"time"
)

// NewDefaultHttpClient 创建默认的http client
func NewDefaultHttpClient() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   5,
			MaxConnsPerHost:       5,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}
