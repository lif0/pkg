package semantic_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/lif0/pkg/semantic"
)

func BenchmarkEstimatePayloadOf(b *testing.B) {
	b.ReportAllocs()

	benchmarkEstimateSizeOf[int](b, func() int { return rand.Intn(32) })
	benchmarkEstimateSizeOf[int8](b, func() int8 { return int8(rand.Intn(32)) })
	benchmarkEstimateSizeOf[int16](b, func() int16 { return int16(rand.Intn(32)) })
	benchmarkEstimateSizeOf[int32](b, func() int32 { return int32(rand.Intn(32)) })
	benchmarkEstimateSizeOf[int64](b, func() int64 { return int64(rand.Intn(32)) })
	benchmarkEstimateSizeOf[uint](b, func() uint { return uint(rand.Intn(32)) })
	benchmarkEstimateSizeOf[uint8](b, func() uint8 { return uint8(rand.Intn(32)) })
	benchmarkEstimateSizeOf[uint16](b, func() uint16 { return uint16(rand.Intn(32)) })
	benchmarkEstimateSizeOf[uint32](b, func() uint32 { return uint32(rand.Intn(32)) })
	benchmarkEstimateSizeOf[uint64](b, func() uint64 { return uint64(32) })
	benchmarkEstimateSizeOf[uintptr](b, func() uintptr { return uintptr(rand.Intn(32)) })
	benchmarkEstimateSizeOf[float32](b, func() float32 { return rand.Float32() })
	benchmarkEstimateSizeOf[float64](b, func() float64 { return rand.Float64() })
	benchmarkEstimateSizeOf[complex64](b, func() complex64 { return complex(1, 2) })
	benchmarkEstimateSizeOf[complex128](b, func() complex128 { return complex(3, 4) })
	benchmarkEstimateSizeOf[bool](b, func() bool { return rand.Intn(32)%2 == 0 })
	benchmarkEstimateSizeOf[string](b, func() string { return "bench" })
	benchmarkEstimateSizeOf[time.Time](b, func() time.Time { return time.Now() })
	benchmarkEstimateSizeOf[time.Duration](b, func() time.Duration { return time.Second })
}

// T, *T
// []T, *[]T, []*T, *[]*T
// [N]T, *[N]T, [N]*T, *[N]*T
// []T{nil}, *[]T{nil}, []*T{nil}, *[]*T{nil}, [N]T{nil}, *[N]T{nil}, [N]*T{nil}, *[N]*T{nil}
func benchmarkEstimateSizeOf[T any](b *testing.B, getValue func() T) {
	// setup T
	val := getValue() // T

	// setup *T
	ptrVal := &val

	// setup []T
	sliceVal := make([]T, 100)
	for i := range sliceVal {
		sliceVal[i] = getValue()
	}
	sliceValAny := sliceVal

	// setup *[]T
	ptrSliceVal := &sliceVal

	// setup []*T
	slicePrtVal := make([]*T, 100)
	for i := range slicePrtVal {
		v := getValue()
		slicePrtVal[i] = &v
	}
	slicePrtValAny := slicePrtVal
	// setup *[]*T
	prtSlicePrtVal := &slicePrtValAny

	// setup [N]T
	arrayVal := [100]T{}
	for i := range arrayVal {
		v := getValue()
		arrayVal[i] = v
	}
	arrayValAny := arrayVal

	// setup  *[N]T
	ptrArrayVal := &arrayValAny

	// setup [N]*T
	arrayPrtVal := make([]*T, 100)
	for i := range arrayPrtVal {
		v := getValue()
		arrayPrtVal[i] = &v
	}
	// setup *[N]*T
	prtArrayPrtVal := &arrayPrtVal

	TName := reflect.TypeOf(val).String()

	// T
	b.Run(TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(val)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})
	// *T
	b.Run("*"+TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(ptrVal)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// []
	b.Run("[]"+TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(sliceValAny)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// *[]T
	b.Run("*[]"+TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(ptrSliceVal)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// []*T
	b.Run("[]*"+TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(slicePrtValAny)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// *[]*T
	b.Run("*[]*"+TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(prtSlicePrtVal)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// [N]T
	b.Run("[100]"+TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(arrayValAny)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// *[N]T
	b.Run("*[100]"+TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(ptrArrayVal)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// [N]*T
	b.Run("[100]*"+TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(arrayPrtVal)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// *[N]*T
	b.Run("*[100]*"+TName, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := semantic.EstimatePayloadOf(prtArrayPrtVal)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// ===== nil cases =====

	// []T{nil}
	b.Run("[]"+TName+"{nil}", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var arg []T
			res := semantic.EstimatePayloadOf(arg)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// *[]T{nil}
	b.Run("*[]"+TName+"{nil}", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var arg *[]T
			res := semantic.EstimatePayloadOf(arg)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// []*T{nil}
	b.Run("[]*"+TName+"{nil}", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var arg []*T
			res := semantic.EstimatePayloadOf(arg)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// *[]*T{nil}
	b.Run("*[]*"+TName+"{nil}", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var arg []*T
			res := semantic.EstimatePayloadOf(&arg)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// [N]T{nil}
	b.Run("[100]"+TName+"{nil}", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var arg [100]T
			res := semantic.EstimatePayloadOf(arg)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// *[N]T{nil}
	b.Run("*[100]"+TName+"{nil}", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var arg *[100]T
			res := semantic.EstimatePayloadOf(arg)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// [N]*T{nil} // it have N nil-objects
	b.Run("[100]*"+TName+"{nil}", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var arg [100]*T
			res := semantic.EstimatePayloadOf(arg)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})

	// *[N]*T{nil} // it have N nil-objects
	b.Run("*[100]*"+TName+"{nil}", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var arg *[100]*T
			res := semantic.EstimatePayloadOf(arg)
			if res == semantic.ErrFailEstimatePayload {
				fmt.Println(TName + " miss switch")
			}
		}
	})
}
