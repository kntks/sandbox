package mutex

import (
	"sync"
	"time"
)

type MyMutex struct {
	mu sync.Mutex
	ma map[int]int
}

func (m *MyMutex) write() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.mu.Lock()
			defer m.mu.Unlock()
			m.ma[i] = i
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}

func (m *MyMutex) read() {
	for i := 0; i < 20; i++ {
		m.mu.Lock()
		defer m.mu.Unlock()
	}
}

type MyRWMutex struct {
	mu sync.RWMutex
	ma map[int]int
}

func (m *MyRWMutex) write() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.mu.Lock()
			defer m.mu.Unlock()
			m.ma[i] = i
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}

func (m *MyRWMutex) read() {
	for i := 0; i < 20; i++ {
		m.mu.RLock()
		defer m.mu.RUnlock()
	}
}
