package sync

import (
	"reflect"
	"strings"
	"sync"
	"testing"
)

func BenchmarkMutexes(b *testing.B) {
	benchmarkMutex(b, func() sync.Locker { return &sync.Mutex{} }, func() { _ = 1 })
	benchmarkMutex(b, func() sync.Locker { return &sync.RWMutex{} }, func() { _ = 1 })
	benchmarkMutex(b, func() sync.Locker { return &ReentrantMutex{} }, func() { _ = 1 })
}

func BenchmarkMutexesParallel(b *testing.B) {
	benchmarkMutexParallel(b, func() sync.Locker { return &sync.Mutex{} }, func() { _ = 1 })
	benchmarkMutexParallel(b, func() sync.Locker { return &sync.RWMutex{} }, func() { _ = 1 })
	benchmarkMutexParallel(b, func() sync.Locker { return &ReentrantMutex{} }, func() { _ = 1 })
}

func benchmarkMutex(b *testing.B, getMutex func() sync.Locker, emulateWork func()) {
	mu := getMutex()
	TName := strings.TrimPrefix(reflect.TypeOf(mu).String(), "*")

	b.Run(TName, func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			mu.Lock()
			emulateWork()
			mu.Unlock()
		}
	})
}

func benchmarkMutexParallel(b *testing.B, getMutex func() sync.Locker, emulateWork func()) {
	mu := getMutex()
	TName := strings.TrimPrefix(reflect.TypeOf(mu).String(), "*")

	b.SetParallelism(100)

	b.Run(TName+"Parallel", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mu.Lock()
				emulateWork()
				mu.Unlock()
			}
		})
	})
}
