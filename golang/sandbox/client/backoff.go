package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/xerrors"
)

type Result struct {
	Status int
	Body   []byte
}

type Backoff struct {
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func RetryDelayGenerator(maxRetry uint) func() int {
	retry := 0
	return func() int {
		defer func() { retry++ }()
		return 1 << min(retry, int(maxRetry))
	}
}

func Get(url string, maxRetry int, retryDelayFunc func() int) (*Result, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	client := http.Client{
		Transport: transport,
	}
	for retry := 0; retry <= maxRetry; retry++ {
		res, err := client.Do(req)
		if err != nil {
			return nil, xerrors.Errorf("failed to Get reqest : %w", err)
		}
		defer res.Body.Close()

		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		switch {
		case http.StatusBadRequest <= res.StatusCode && res.StatusCode < http.StatusInternalServerError:
			return &Result{
				Status: res.StatusCode,
				Body:   bytes,
			}, xerrors.Errorf("code: %d, request: %+v", res.StatusCode, *res.Request)
		case http.StatusInternalServerError <= res.StatusCode:
			if retry != maxRetry {
				time.Sleep(time.Duration(retryDelayFunc()) * time.Second)
				continue
			}
			return &Result{
				Status: res.StatusCode,
				Body:   bytes,
			}, xerrors.Errorf("code: %d, request: %+v", res.StatusCode, *res.Request)
		}
		return &Result{
			Status: res.StatusCode,
			Body:   bytes,
		}, nil
	}
	return nil, xerrors.New("failed to Get request")
}

func SampleGet(url string, result *Result) func() error {
	retry := 0
	return func() error {
		fmt.Println(retry)
		retry++
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		transport := http.DefaultTransport.(*http.Transport).Clone()
		client := http.Client{
			Transport: transport,
		}
		res, err := client.Do(req)
		if err != nil {
			return xerrors.Errorf("failed to Get reqest : %w", err)
		}
		defer res.Body.Close()

		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		switch {
		case http.StatusBadRequest <= res.StatusCode && res.StatusCode < http.StatusInternalServerError:
			result.Status = res.StatusCode
			result.Body = bytes
			return xerrors.Errorf("code: %d, request: %+v", res.StatusCode, *res.Request)
		case http.StatusInternalServerError <= res.StatusCode:
			result.Status = res.StatusCode
			result.Body = bytes
			return xerrors.Errorf("code: %d, request: %+v", res.StatusCode, *res.Request)
		}
		return nil
	}
}

func ExampleExponential(url string, maxRetry uint64) {
	b := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetry)
	var result Result
	err := backoff.Retry(SampleGet(url, &result), b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}

func shouldRetry(status int) bool {
	if http.StatusBadRequest <= status && status < http.StatusInternalServerError {
		return false
	}

	return http.StatusInternalServerError <= status
}

// func Retry(ctx context.Context, f func() (*http.Response, error), backoff BackoffStrategy) (*http.Response, error) {
// 	for {
// 		resp, err := f()

// 		if resp != nil && !shouldRetry(resp.StatusCode) {
// 			return resp, err
// 		}
// 		// Return if we shouldn't retry.
// 		pause, retry := backoff.Pause()
// 		if !retry {
// 			return resp, err
// 		}

// 		if resp.Body != nil {
// 			resp.Body.Close()
// 		}

// 		var done <-chan struct{}
// 		if ctx != nil {
// 			done = ctx.Done()
// 		}
// 		select {
// 		case <-done:
// 			return nil, ctx.Err()
// 		case <-time.After(pause):
// 		}
// 	}
// }
