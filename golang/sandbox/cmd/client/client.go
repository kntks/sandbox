package main

import (
	"fmt"
	"log"
	"sandbox/client"
	"sync"
	"time"
)

func httpStats() {
	wg := sync.WaitGroup{}
	client := client.NewClient()
	loopNum := 5000

	wg.Add(loopNum)
	// go run cmd/server/main.go
	url := "http://localhost:8081/200"
	start := time.Now()
	for i := 0; i < loopNum; i++ {
		go func(x int) {
			defer wg.Done()
			client.GetAndDiscard(url)
			fmt.Printf("====== %d回 =======\n", x+1)
			fmt.Printf("%+v\n", client.StatResult())
		}(i)
	}
	wg.Wait()
	fmt.Printf("%f秒\n", time.Since(start).Seconds())
}

func customExponentialBackoff() {
	url := "http://localhost:8081/5xx"
	maxRetry := 3
	result, err := client.Get(url, maxRetry, client.RetryDelayGenerator(uint(maxRetry)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" %+v \n", string(result.Body))
}

func main() {
	url := "http://localhost:8081/4xx"
	client.ExampleExponential(url, 4)
}
