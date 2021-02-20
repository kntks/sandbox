package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{OK}`))
})

func TestLoopOutClient(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	loopOutClient(ts.URL)
}

func TestLoopInClient(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	loopInClient(ts.URL)
}

// case1-1: forループの外でclientを作成
func BenchmarkCase11(b *testing.B) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loopOutClient(ts.URL)
	}
}

// case1-2: forループの中でclientを作成
func BenchmarkCase12(b *testing.B) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loopInClient(ts.URL)
	}
}
