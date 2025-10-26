package concurrency_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	concurrency "github.com/lif0/pkg/concurrency"
)

func benchReadHeavy(b *testing.B, readFn func(), writeFn func()) {
	// ~99% read, 1% write
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%100 == 0 {
				writeFn()
			} else {
				readFn()
			}
			i++
		}
	})
}

func benchWriteHeavy(b *testing.B, writeFn func()) {
	// 100% write
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			writeFn()
		}
	})
}

func benchMixed(b *testing.B, readFn func(), writeFn func()) {
	// ~80% read / 20% write
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%20 == 0 {
				writeFn()
			} else {
				readFn()
			}
			i++
		}
	})
}

func Benchmark_Int64_ReadHeavy(b *testing.B) {
	b.Run("SyncValue", func(b *testing.B) {
		sv := concurrency.NewSyncValue[int64]()
		benchReadHeavy(b,
			func() {
				sv.ReadValue(func(v *int64) { _ = *v })
			},
			func() {
				sv.MutateValue(func(v *int64) { *v++ })
			},
		)
	})
	b.Run("Atomic", func(b *testing.B) {
		var v atomic.Int64
		benchReadHeavy(b,
			func() { _ = v.Load() },
			func() { v.Add(1) },
		)
	})
}

func Benchmark_Int64_WriteHeavy(b *testing.B) {
	b.Run("SyncValue", func(b *testing.B) {
		sv := concurrency.NewSyncValue[int64]()
		benchWriteHeavy(b, func() { sv.MutateValue(func(v *int64) { *v++ }) })
	})
	b.Run("Atomic", func(b *testing.B) {
		var v atomic.Int64
		benchWriteHeavy(b, func() { v.Add(1) })
	})
}

func Benchmark_Int64_Mixed(b *testing.B) {
	b.Run("SyncValue", func(b *testing.B) {
		sv := concurrency.NewSyncValue[int64]()
		benchMixed(b,
			func() {
				sv.ReadValue(func(v *int64) { _ = *v })
			},
			func() {
				sv.MutateValue(func(v *int64) { *v++ })
			},
		)
	})
	b.Run("Atomic", func(b *testing.B) {
		var v atomic.Int64
		benchMixed(b,
			func() { _ = v.Load() },
			func() { v.Add(1) },
		)
	})
}

func Benchmark_Map_Mixed(b *testing.B) {
	b.Run("SyncValue", func(b *testing.B) {
		sv := concurrency.NewSyncValue[map[string]int](map[string]int{"x": 0})
		benchMixed(b,
			func() {
				sv.ReadValue(func(v *map[string]int) { _ = (*v)["x"] })
			},
			func() {
				sv.MutateValue(func(v *map[string]int) { (*v)["x"]++ })
			},
		)
	})

	b.Run("sync.Map", func(b *testing.B) {
		sv := sync.Map{}
		sv.Store("x", int(0))

		benchMixed(b,
			func() {
				sv.Load("x")
			},
			func() {
				v, _ := sv.Load("x")
				vv := v.(int) + 1
				sv.Store("x", vv)
			},
		)
	})
}

func Benchmark_Complex_Mixed(b *testing.B) {
	b.Run("SyncValue", func(b *testing.B) {
		sv := concurrency.NewSyncValue[complexStruct](complexStruct{Flag: true, Nums: []int{1}, Index: map[string]int{"x": 1}})
		benchMixed(b,
			func() {
				sv.ReadValue(func(v *complexStruct) {
					_ = v.Flag
					_ = v.Index["x"]
					_ = v.Nums[0]
				})
			},
			func() {
				sv.MutateValue(func(v *complexStruct) {
					v.Flag = !v.Flag

					for i := 0; i < 10; i++ {
						v.Nums = append(v.Nums, i)
					}
					for i := 0; i < 10; i++ {
						v.Index[fmt.Sprintf("k%d", i)]++
					}
				})
			},
		)
	})
	b.Run("Atomic.Pointer", func(b *testing.B) {
		var v atomic.Value
		cs := complexStruct{Flag: true, Nums: []int{1}, Index: map[string]int{"x": 1}}
		v.Store(cs)
		benchMixed(b,
			func() {
				cva := v.Load()
				cv := cva.(complexStruct)
				_ = cv.Flag
				_ = cv.Index["x"]
				_ = cv.Nums[0]
			},
			func() {
				cva := v.Load()
				old := cva.(complexStruct)

				new := complexStruct{}

				new.Flag = old.Flag
				new.Nums = make([]int, len(old.Nums))
				copy(new.Nums, old.Nums)
				new.Index = map[string]int{}
				for k, v := range old.Index {
					new.Index[k] = v
				}

				//  modify

				new.Flag = !new.Flag
				for i := 0; i < 10; i++ {
					new.Nums = append(new.Nums, i)
				}
				for i := 0; i < 10; i++ {
					new.Index[fmt.Sprintf("k%d", i)]++
				}

				v.Store(new)
			},
		)
	})
}
