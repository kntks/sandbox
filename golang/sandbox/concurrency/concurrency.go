package concurrency

import (
	"runtime"
	"sync"
	"time"
)

// start := time.Now()
// end := time.Now()
// fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

// forの中にgo routineを10個立てる
// wait groupのみで制御
func Case1() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Microsecond)
		}()
	}
	wg.Wait()
}

// case1と同じだがgo routineの数をCPUコア数で制御する
func Case2() {
	wg := sync.WaitGroup{}
	cpus := runtime.NumCPU() // CPUの数
	semaphore := make(chan int, cpus)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- 1
			time.Sleep(100 * time.Millisecond)
			<-semaphore
		}()
	}
	wg.Wait()
}

func Case3() {
	done := make(chan bool)
	go func() {
		defer close(done)
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
		}
	}()
	<-done
}
