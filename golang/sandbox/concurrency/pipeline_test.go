package concurrency

import "testing"

func BenchmarkSliceProcess(b *testing.B) {
	SliceProcess()
}

func BenchmarkPipeline(b *testing.B) {
	PipelineProcess()
}
