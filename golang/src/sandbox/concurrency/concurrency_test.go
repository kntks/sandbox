package concurrency

import "testing"

func TestCase1(t *testing.T) {
	Case1()
}

func TestCase2(t *testing.T) {
	Case2()
} 

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
