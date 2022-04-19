package main

import (
	"context"
	"fmt"
	"sandbox/concurrency"
	"time"
)

func main() {
	// concurrency.Case3()
	// concurrency.LockMap()
	// PipelineMain()
	GeneratorMain()
}

// http://tmrts.com/go-patterns/concurrency/generator.html
// https://blog.web-apps.tech/generator-pattern/
func GeneratorMain() {
	for v := range generators(context.Background()) {
		fmt.Println("処理中", v)
	}
}
func generators(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				time.Sleep(time.Duration(i) * time.Second)
			}
		}
	}()
	return ch
}

func PipelineMain() {
	start := time.Now()
	concurrency.SliceProcess()
	fmt.Println("sliece: ", time.Since(start))

	start = time.Now()
	concurrency.PipelineProcess()
	fmt.Println("pipeline: ", time.Since(start))
}
