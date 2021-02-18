package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func LockMap() {
	var wg sync.WaitGroup
	var mu sync.RWMutex

	rand.Seed(time.Now().UnixNano())
	m := make(map[int]int, 10)
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}(i)

		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i * 2
			mu.Unlock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}(i)
	}
	wg.Wait()

	fmt.Println(m)
}
