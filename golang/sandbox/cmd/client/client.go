package main

import (
	"fmt"
	"sandbox/client"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	client := client.NewClient()
	loopNum := 5000

	wg.Add(loopNum)
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
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
