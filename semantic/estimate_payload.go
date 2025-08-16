package semantic

import (
	"reflect"
	"time"
	"unsafe"
)

var ErrFailEstimatePayload int = -1

// EstimatePayloadOf returns an approximate payload size (in bytes) of the given value.
//
// This function performs **zero allocations** and runs with **0 B/op**.
//
// Supported types:
//   - all scalar types and their pointers:
//     int, int8, int16, int32 (rune), int64,
//     uint, uint8 (byte), uint16, uint32, uint64, uintptr,
//     float32, float64,
//     complex64, complex128,
//     bool,
//     string,
//     time.Time, time.Duration
//   - slices of the above types and their pointers:
//     []T, *[]T, []*T, *[]*T
//   - arrays [N]T and *[N]T (via reflection)
//
// For pointer and slice values, nil is treated as zero size.
// For strings and slices of strings, the sum of actual string lengths is returned.
//
// If the type is not recognized or cannot be handled, the function
// returns ErrFailEstimatePayload.
//
// Performance note:
//
//	For arrays [N]T it is highly recommended to pass a pointer (*[N]T)
//	instead of the array value itself. Passing an array by value causes the
//	entire array to be copied when calling this function, which can make
//	estimation dozens or even hundreds of times slower for large arrays.
func EstimatePayloadOf(arg any) int {
	switch v := arg.(type) {
	case nil:
		return 0

	// ----- scalar T / *T -----
	case int:
		return int(unsafe.Sizeof(v)) // depends on the architecture
	case *int:
		if v == nil {
			return 0
		}
		return int(unsafe.Sizeof(*v)) // depends on the architecture
	case int8:
		return 1
	case *int8:
		if v == nil {
			return 0
		}
		return 1
	case int16:
		return 2
	case *int16:
		if v == nil {
			return 0
		}
		return 2
	case int32: // rune
		return 4
	case *int32:
		if v == nil {
			return 0
		}
		return 4
	case int64:
		return 8
	case *int64:
		if v == nil {
			return 0
		}
		return 8

	case uint:
		return int(unsafe.Sizeof(v)) // depends on the architecture
	case *uint:
		if v == nil {
			return 0
		}
		return int(unsafe.Sizeof(*v)) // depends on the architecture
	case uint8:
		return 1
	case *uint8: // byte
		if v == nil {
			return 0
		}
		return 1
	case uint16:
		return 2
	case *uint16:
		if v == nil {
			return 0
		}
		return 2
	case uint32:
		return 4
	case *uint32:
		if v == nil {
			return 0
		}
		return 4
	case uint64:
		return 8
	case *uint64:
		if v == nil {
			return 0
		}
		return 8
	case uintptr:
		return int(unsafe.Sizeof(v))
	case *uintptr:
		if v == nil {
			return 0
		}
		return int(unsafe.Sizeof(*v))

	case float32:
		return 4
	case *float32:
		if v == nil {
			return 0
		}
		return 4
	case float64:
		return 8
	case *float64:
		if v == nil {
			return 0
		}
		return 8

	case complex64:
		return 8
	case *complex64:
		if v == nil {
			return 0
		}
		return 8
	case complex128:
		return 16
	case *complex128:
		if v == nil {
			return 0
		}
		return 16

	case bool:
		return 1
	case *bool:
		if v == nil {
			return 0
		}
		return 1

	case string:
		return len(v)
	case *string:
		if v == nil {
			return 0
		}
		return len(*v)

	case time.Time:
		return int(unsafe.Sizeof(v)) // depends on the architecture
	case *time.Time:
		if v == nil {
			return 0
		}
		return int(unsafe.Sizeof(*v)) // depends on the architecture

	case time.Duration:
		return 8
	case *time.Duration:
		if v == nil {
			return 0
		}
		return 8

	// ----- []T / *[]T -----
	case []int:
		return len(v) * int(unsafe.Sizeof(int(0)))
	case *[]int:
		if v == nil {
			return 0
		}
		return len(*v) * int(unsafe.Sizeof(int(0)))
	case []*int:
		return countNonNilPtrs(v) * int(unsafe.Sizeof(int(0)))
	case *[]*int:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * int(unsafe.Sizeof(int(0)))

	case []int8:
		return len(v)
	case *[]int8:
		if v == nil {
			return 0
		}
		return len(*v)
	case []*int8:
		return countNonNilPtrs(v) * 1
	case *[]*int8:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 1

	case []int16:
		return len(v) * 2
	case *[]int16:
		if v == nil {
			return 0
		}
		return len(*v) * 2
	case []*int16:
		return countNonNilPtrs(v) * 2
	case *[]*int16:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 2

	case []int32:
		return len(v) * 4
	case *[]int32:
		if v == nil {
			return 0
		}
		return len(*v) * 4
	case []*int32:
		return countNonNilPtrs(v) * 4
	case *[]*int32:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 4

	case []int64:
		return len(v) * 8
	case *[]int64:
		if v == nil {
			return 0
		}
		return len(*v) * 8
	case []*int64:
		return countNonNilPtrs(v) * 8
	case *[]*int64:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 8

	case []uint:
		return len(v) * int(unsafe.Sizeof(uint(0)))
	case *[]uint:
		if v == nil {
			return 0
		}
		return len(*v) * int(unsafe.Sizeof(uint(0)))
	case []*uint:
		return countNonNilPtrs(v) * int(unsafe.Sizeof(uint(0)))
	case *[]*uint:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * int(unsafe.Sizeof(uint(0)))

	case []uint8:
		return len(v)
	case *[]uint8:
		if v == nil {
			return 0
		}
		return len(*v)
	case []*uint8:
		return countNonNilPtrs(v) * 1
	case *[]*uint8:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 1

	case []uint16:
		return len(v) * 2
	case *[]uint16:
		if v == nil {
			return 0
		}
		return len(*v) * 2
	case []*uint16:
		return countNonNilPtrs(v) * 2
	case *[]*uint16:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 2

	case []uint32:
		return len(v) * 4
	case *[]uint32:
		if v == nil {
			return 0
		}
		return len(*v) * 4
	case []*uint32:
		return countNonNilPtrs(v) * 4
	case *[]*uint32:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 4

	case []uint64:
		return len(v) * 8
	case *[]uint64:
		if v == nil {
			return 0
		}
		return len(*v) * 8
	case []*uint64:
		return countNonNilPtrs(v) * 8
	case *[]*uint64:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 8

	case []uintptr:
		return len(v) * int(unsafe.Sizeof(uintptr(0)))
	case *[]uintptr:
		if v == nil {
			return 0
		}
		return len(*v) * int(unsafe.Sizeof(uintptr(0)))
	case []*uintptr:
		return countNonNilPtrs(v) * int(unsafe.Sizeof(uintptr(0)))
	case *[]*uintptr:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * int(unsafe.Sizeof(uintptr(0)))

	case []float32:
		return len(v) * 4
	case *[]float32:
		if v == nil {
			return 0
		}
		return len(*v) * 4
	case []*float32:
		return countNonNilPtrs(v) * 4
	case *[]*float32:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 4

	case []float64:
		return len(v) * 8
	case *[]float64:
		if v == nil {
			return 0
		}
		return len(*v) * 8
	case []*float64:
		return countNonNilPtrs(v) * 8
	case *[]*float64:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 8

	case []complex64:
		return len(v) * 8
	case *[]complex64:
		if v == nil {
			return 0
		}
		return len(*v) * 8
	case []*complex64:
		return countNonNilPtrs(v) * 8
	case *[]*complex64:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 8

	case []complex128:
		return len(v) * 16
	case *[]complex128:
		if v == nil {
			return 0
		}
		return len(*v) * 16
	case []*complex128:
		return countNonNilPtrs(v) * 16
	case *[]*complex128:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 16

	case []bool:
		return len(v)
	case *[]bool:
		if v == nil {
			return 0
		}
		return len(*v)
	case []*bool:
		return countNonNilPtrs(v) * 1
	case *[]*bool:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v)

	case []string:
		size := 0
		for _, s := range v {
			size += len(s)
		}
		return size
	case *[]string:
		if v == nil {
			return 0
		}
		size := 0
		for _, s := range *v {
			size += len(s)
		}
		return size
	case []*string:
		return sumNonNilPtrsLenString(v)
	case *[]*string:
		if v == nil {
			return 0
		}
		return sumNonNilPtrsLenString(*v)

	case []time.Time:
		return len(v) * int(unsafe.Sizeof(time.Time{}))
	case *[]time.Time:
		if v == nil {
			return 0
		}
		return len(*v) * int(unsafe.Sizeof(time.Time{}))
	case []*time.Time:
		return countNonNilPtrs(v) * int(unsafe.Sizeof(time.Time{}))
	case *[]*time.Time:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * int(unsafe.Sizeof(time.Time{}))

	case []time.Duration:
		return len(v) * 8 // duration alias of int64
	case *[]time.Duration:
		if v == nil {
			return 0
		}
		return len(*v) * 8 // duration alias of int64
	case []*time.Duration:
		return countNonNilPtrs(v) * 8 // duration alias of int64
	case *[]*time.Duration:
		if v == nil {
			return 0
		}
		return countNonNilPtrs(*v) * 8 // duration alias of int64

	// ----- fallback: [N]T / *[N]T -----
	default:
		rv := reflect.ValueOf(arg)
		if !rv.IsValid() {
			return 0
		}
		// dereferenced pointers; for *[N]T{nil} return 0
		for rv.Kind() == reflect.Ptr {
			if rv.IsNil() {
				return 0
			}
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.Array {
			return payloadArray(rv)
		}
		return ErrFailEstimatePayload
	}
}

func countNonNilPtrs[T any](s []*T) (c int) {
	c = 0
	for _, p := range s {
		if p != nil {
			c++
		}
	}
	return c
}

func sumNonNilPtrsLenString(s []*string) (size int) {
	size = 0
	for _, p := range s {
		if p != nil {
			size += len(*p)
		}
	}
	return size
}

func payloadArray(rv reflect.Value) int {
	// rv — must be dereferenced pointer already
	n := rv.Len()
	et := rv.Type().Elem()

	switch et.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Bool:
		return n * int(et.Size())

	case reflect.String:
		t := 0
		for i := 0; i < n; i++ {
			t += rv.Index(i).Len() // payload srt
		}
		return t

	case reflect.Struct:
		// Freq fast-path: time.Time
		if et == reflect.TypeOf(time.Time{}) {
			return n * int(et.Size())
		}
		return ErrFailEstimatePayload // unknown struct - return fail

	case reflect.Ptr:
		el := et.Elem()
		switch el.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128,
			reflect.Bool:
			sz := int(el.Size())
			total := 0
			for i := 0; i < n; i++ {
				if !rv.Index(i).IsNil() {
					total += sz
				}
			}
			return total
		case reflect.String:
			total := 0
			for i := 0; i < n; i++ {
				x := rv.Index(i)
				if !x.IsNil() {
					total += x.Elem().Len()
				}
			}
			return total
		case reflect.Struct:
			if el == reflect.TypeOf(time.Time{}) {
				sz := int(el.Size())
				total := 0
				for i := 0; i < n; i++ {
					if !rv.Index(i).IsNil() {
						total += sz
					}
				}
				return total
			}
			return ErrFailEstimatePayload
		default:
			return ErrFailEstimatePayload
		}
	default:
		return ErrFailEstimatePayload
	}
}
