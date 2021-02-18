package concurrency

import "testing"

func BenchmarkCase1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Case1()
	}
}
func BenchmarkCase2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Case2()
	}
}

func BenchmarkCase3(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Case3()
	}
}
