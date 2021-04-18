package main

import (
	"fmt"
	"sandbox/concurrency"
	"time"
)

func main() {
	// concurrency.Case3()
	concurrency.LockMap()
	PipelineMain()
}

func PipelineMain() {
	start := time.Now()
	concurrency.SliceProcess()
	fmt.Println("sliece: ", time.Since(start))

	start = time.Now()
	concurrency.PipelineProcess()
	fmt.Println("pipeline: ", time.Since(start))
}
