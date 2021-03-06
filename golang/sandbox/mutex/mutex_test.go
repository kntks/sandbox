package mutex

import (
	"sync"
	"testing"
)

func BenchmarkMyMutex(t *testing.B) {
	m := MyMutex{mu: sync.Mutex{}, ma: map[int]int{}}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		go m.read()
	}
}

func BenchmarkMyRWMutex(t *testing.B) {
	m := MyRWMutex{mu: sync.RWMutex{}, ma: map[int]int{}}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		go m.read()
	}
}
