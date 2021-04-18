package concurrency

import (
	"time"
)

func processA1() []struct{} {
	list := []struct{}{{}, {}, {}, {}, {}}
	for range list {
		time.Sleep(100 * time.Millisecond)
	}
	return list
}

func processA2(data []struct{}) []struct{} {
	for range data {
		time.Sleep(200 * time.Millisecond)
	}
	return data
}

func SliceProcess() {
	processA2(processA1())
}

func processB1() <-chan struct{} {
	ch := make(chan struct{})
	list := []struct{}{{}, {}, {}, {}, {}}
	go func() {
		defer close(ch)
		for _, x := range list {
			time.Sleep(100 * time.Millisecond)
			ch <- x
		}
	}()
	return ch
}

func processB2(in <-chan struct{}) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		for x := range in {
			time.Sleep(200 * time.Millisecond)
			ch <- x
		}
	}()
	return ch
}

func PipelineProcess() {
	processB2(processB1())
}
