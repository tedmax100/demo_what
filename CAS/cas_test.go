package cas

import (
	"sync"
	"sync/atomic"
	"testing"
)

var casValue int64
var faaValue int64

func BenchmarkMutexCAS(b *testing.B) {
	var mu sync.Mutex
	casValue = 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		if casValue == int64(i-1) {
			casValue = int64(i)
		}
		mu.Unlock()
	}
}

func BenchmarkAtomicCAS(b *testing.B) {
	casValue = 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		old := int64(i - 1)
		new := int64(i)
		for !atomic.CompareAndSwapInt64(&casValue, old, new) {
			old = casValue
		}
	}

}

func BenchmarkMutexAdd(b *testing.B) {
	var mu sync.Mutex
	faaValue = 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		faaValue++
		mu.Unlock()
	}
}

func BenchmarkAtomicAdd(b *testing.B) {
	faaValue = 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&faaValue, 1)
	}
}
