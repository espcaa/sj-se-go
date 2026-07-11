package sj

import (
	"net/http"
	"time"
)

type Client struct {
	baseUrl         string
	httpClient      *http.Client
	subscriptionKey string
}

func NewClient(baseUrl, subscriptionKey string) *Client {
	return &Client{
		baseUrl: baseUrl,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		subscriptionKey: subscriptionKey,
	}
}
