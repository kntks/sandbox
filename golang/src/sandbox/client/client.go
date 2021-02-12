package client

import (
	"net/http"
	"time"
)

type Client struct {
	c *http.Client
}

func NewClient(timeout int) *Client {
	return &Client{
		c: &http.Client{
			Timeout: time.Duration(timeout) * time.Minute,
		},
	}
}

func (c *Client) get() {}
