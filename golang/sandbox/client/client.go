package client

import (
	"log"
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

func loopOutClient(url string) {
	client := &http.Client{}
	for i := 0; i < 10; i++ {
		_, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loopInClinet(url string) {
	for i := 0; i < 10; i++ {
		client := &http.Client{}
		_, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
	}
}
