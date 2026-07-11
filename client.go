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

type ClientOption func(*Client)

func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = hc
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

func NewClient(baseUrl, subscriptionKey string, opts ...ClientOption) *Client {
	return &Client{
		baseUrl: baseUrl,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		subscriptionKey: subscriptionKey,
	}
}
