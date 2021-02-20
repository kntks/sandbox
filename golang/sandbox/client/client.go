package client

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/tcnksm/go-httpstat"
)

type Client struct {
	*http.Client
	stat httpstat.Result
}

func NewClient() *Client {
	return &Client{
		&http.Client{
			Transport: CustomTransport(),
		},
		httpstat.Result{},
	}
}

func (c *Client) GetAndDiscard(url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		return err
	}
	result.End(time.Now())
	c.stat = result
	return nil
}

func (c *Client) StatResult() httpstat.Result {
	return c.stat
}

func CustomTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   2,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func loopOutClient(url string) {
	wg := sync.WaitGroup{}
	client := &http.Client{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := client.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			io.Copy(ioutil.Discard, res.Body)
		}()
	}
	wg.Wait()
}

func loopInClient(url string) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{}
			res, err := client.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			io.Copy(ioutil.Discard, res.Body)
		}()
	}
	wg.Wait()
}
