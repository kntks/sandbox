package main

import (
	"fmt"
	"time"
)

func ticker() {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	done := make(chan struct{})
	defer close(done)

	go func() {
		for {
			select {
			case now := <-t.C:
				fmt.Println(now.Format(time.RFC3339))
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	done <- struct{}{}
}

func main() {
	ticker()
}
