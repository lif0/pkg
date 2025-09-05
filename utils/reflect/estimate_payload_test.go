package reflect_test

import (
	"testing"
	"time"
	"unsafe"

	"github.com/lif0/pkg/utils/reflect"
)

var (
	szInt        = int(unsafe.Sizeof(int(0)))
	szInt8       = int(unsafe.Sizeof(int8(0)))
	szInt16      = int(unsafe.Sizeof(int16(0)))
	szInt32      = int(unsafe.Sizeof(int32(0)))
	szInt64      = int(unsafe.Sizeof(int64(0)))
	szUint       = int(unsafe.Sizeof(uint(0)))
	szUint8      = int(unsafe.Sizeof(uint8(0)))
	szUint16     = int(unsafe.Sizeof(uint16(0)))
	szUint32     = int(unsafe.Sizeof(uint32(0)))
	szUint64     = int(unsafe.Sizeof(uint64(0)))
	szUintptr    = int(unsafe.Sizeof(uintptr(0)))
	szFloat32    = int(unsafe.Sizeof(float32(0)))
	szFloat64    = int(unsafe.Sizeof(float64(0)))
	szComplex64  = int(unsafe.Sizeof(complex64(0)))
	szComplex128 = int(unsafe.Sizeof(complex128(0)))
	szBool       = int(unsafe.Sizeof(false))
	szTime       = int(unsafe.Sizeof(time.Time{}))
	szDuration   = int(unsafe.Sizeof(time.Duration(0)))
)

// Define a custom struct for error cases
type CustomStruct struct{}

func TestEstimatePayloadOf_nil(t *testing.T) {
	got := reflect.EstimatePayloadOf(nil)
	want := 0
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestEstimatePayloadOf_int(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		v := int(42)
		got := reflect.EstimatePayloadOf(v)
		want := szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int", func(t *testing.T) {
		v := int(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int_nil", func(t *testing.T) {
		var pv *int = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int", func(t *testing.T) {
		v := []int{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int empty", func(t *testing.T) {
		v := []int{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int nil", func(t *testing.T) {
		var v []int = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int non-nil", func(t *testing.T) {
		s := []int{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int nil", func(t *testing.T) {
		var ps *[]int = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int", func(t *testing.T) {
		a, b, c := int(1), int(2), int(3)
		v := []*int{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int with nil", func(t *testing.T) {
		a, b := int(1), int(2)
		v := []*int{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int nil", func(t *testing.T) {
		var v []*int = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int all nil", func(t *testing.T) {
		v := []*int{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int non-nil", func(t *testing.T) {
		a, b := int(1), int(2)
		s := []*int{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int nil", func(t *testing.T) {
		var ps *[]*int = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]int", func(t *testing.T) {
		v := [3]int{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int non-nil", func(t *testing.T) {
		v := [3]int{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int nil", func(t *testing.T) {
		var pv *[3]int = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int", func(t *testing.T) {
		a, b, c := int(1), int(2), int(3)
		v := [3]*int{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int with nil", func(t *testing.T) {
		a, b := int(1), int(2)
		v := [3]*int{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int all nil", func(t *testing.T) {
		v := [3]*int{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int non-nil", func(t *testing.T) {
		a, b := int(1), int(2)
		v := [3]*int{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int nil", func(t *testing.T) {
		var pv *[3]*int = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]int", func(t *testing.T) {
		var v [100]int
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szInt
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]int nil", func(t *testing.T) {
		var pv *[100]int = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*int all nil", func(t *testing.T) {
		var v [100]*int
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*int nil", func(t *testing.T) {
		var pv *[100]*int = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_int8(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		v := int8(42)
		got := reflect.EstimatePayloadOf(v)
		want := szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int8 non-nil", func(t *testing.T) {
		v := int8(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int8 nil", func(t *testing.T) {
		var pv *int8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int8", func(t *testing.T) {
		v := []int8{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int8 empty", func(t *testing.T) {
		v := []int8{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int8 nil", func(t *testing.T) {
		var v []int8 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int8 non-nil", func(t *testing.T) {
		s := []int8{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int8 nil", func(t *testing.T) {
		var ps *[]int8 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int8", func(t *testing.T) {
		a, b, c := int8(1), int8(2), int8(3)
		v := []*int8{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int8 with nil", func(t *testing.T) {
		a, b := int8(1), int8(2)
		v := []*int8{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int8 all nil", func(t *testing.T) {
		v := []*int8{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int8 non-nil", func(t *testing.T) {
		a, b := int8(1), int8(2)
		s := []*int8{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int8 nil", func(t *testing.T) {
		var ps *[]*int8 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]int8", func(t *testing.T) {
		v := [3]int8{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int8 non-nil", func(t *testing.T) {
		v := [3]int8{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int8 nil", func(t *testing.T) {
		var pv *[3]int8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int8", func(t *testing.T) {
		a, b, c := int8(1), int8(2), int8(3)
		v := [3]*int8{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int8 with nil", func(t *testing.T) {
		a, b := int8(1), int8(2)
		v := [3]*int8{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int8 all nil", func(t *testing.T) {
		v := [3]*int8{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int8 non-nil", func(t *testing.T) {
		a, b := int8(1), int8(2)
		v := [3]*int8{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int8 nil", func(t *testing.T) {
		var pv *[3]*int8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]int8", func(t *testing.T) {
		var v [100]int8
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szInt8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]int8 nil", func(t *testing.T) {
		var pv *[100]int8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*int8 all nil", func(t *testing.T) {
		var v [100]*int8
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*int8 nil", func(t *testing.T) {
		var pv *[100]*int8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_int16(t *testing.T) {
	t.Run("int16", func(t *testing.T) {
		v := int16(42)
		got := reflect.EstimatePayloadOf(v)
		want := szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int16 non-nil", func(t *testing.T) {
		v := int16(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int16 nil", func(t *testing.T) {
		var pv *int16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int16", func(t *testing.T) {
		v := []int16{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int16 empty", func(t *testing.T) {
		v := []int16{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int16 nil", func(t *testing.T) {
		var v []int16 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int16 non-nil", func(t *testing.T) {
		s := []int16{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int16 nil", func(t *testing.T) {
		var ps *[]int16 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int16", func(t *testing.T) {
		a, b, c := int16(1), int16(2), int16(3)
		v := []*int16{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int16 with nil", func(t *testing.T) {
		a, b := int16(1), int16(2)
		v := []*int16{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int16 all nil", func(t *testing.T) {
		v := []*int16{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int16 non-nil", func(t *testing.T) {
		a, b := int16(1), int16(2)
		s := []*int16{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int16 nil", func(t *testing.T) {
		var ps *[]*int16 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]int16", func(t *testing.T) {
		v := [3]int16{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int16 non-nil", func(t *testing.T) {
		v := [3]int16{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int16 nil", func(t *testing.T) {
		var pv *[3]int16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int16", func(t *testing.T) {
		a, b, c := int16(1), int16(2), int16(3)
		v := [3]*int16{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int16 with nil", func(t *testing.T) {
		a, b := int16(1), int16(2)
		v := [3]*int16{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int16 all nil", func(t *testing.T) {
		v := [3]*int16{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int16 non-nil", func(t *testing.T) {
		a, b := int16(1), int16(2)
		v := [3]*int16{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int16 nil", func(t *testing.T) {
		var pv *[3]*int16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]int16", func(t *testing.T) {
		var v [100]int16
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szInt16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]int16 nil", func(t *testing.T) {
		var pv *[100]int16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*int16 all nil", func(t *testing.T) {
		var v [100]*int16
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*int16 nil", func(t *testing.T) {
		var pv *[100]*int16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_int32(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		v := int32(42)
		got := reflect.EstimatePayloadOf(v)
		want := szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int32 non-nil", func(t *testing.T) {
		v := int32(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int32 nil", func(t *testing.T) {
		var pv *int32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int32", func(t *testing.T) {
		v := []int32{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int32 empty", func(t *testing.T) {
		v := []int32{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int32 nil", func(t *testing.T) {
		var v []int32 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int32 non-nil", func(t *testing.T) {
		s := []int32{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int32 nil", func(t *testing.T) {
		var ps *[]int32 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int32", func(t *testing.T) {
		a, b, c := int32(1), int32(2), int32(3)
		v := []*int32{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int32 with nil", func(t *testing.T) {
		a, b := int32(1), int32(2)
		v := []*int32{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int32 all nil", func(t *testing.T) {
		v := []*int32{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int32 non-nil", func(t *testing.T) {
		a, b := int32(1), int32(2)
		s := []*int32{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int32 nil", func(t *testing.T) {
		var ps *[]*int32 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]int32", func(t *testing.T) {
		v := [3]int32{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int32 non-nil", func(t *testing.T) {
		v := [3]int32{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int32 nil", func(t *testing.T) {
		var pv *[3]int32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int32", func(t *testing.T) {
		a, b, c := int32(1), int32(2), int32(3)
		v := [3]*int32{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int32 with nil", func(t *testing.T) {
		a, b := int32(1), int32(2)
		v := [3]*int32{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int32 all nil", func(t *testing.T) {
		v := [3]*int32{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int32 non-nil", func(t *testing.T) {
		a, b := int32(1), int32(2)
		v := [3]*int32{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int32 nil", func(t *testing.T) {
		var pv *[3]*int32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]int32", func(t *testing.T) {
		var v [100]int32
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szInt32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]int32 nil", func(t *testing.T) {
		var pv *[100]int32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*int32 all nil", func(t *testing.T) {
		var v [100]*int32
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*int32 nil", func(t *testing.T) {
		var pv *[100]*int32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_int64(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		v := int64(42)
		got := reflect.EstimatePayloadOf(v)
		want := szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int64 non-nil", func(t *testing.T) {
		v := int64(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*int64 nil", func(t *testing.T) {
		var pv *int64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int64", func(t *testing.T) {
		v := []int64{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int64 empty", func(t *testing.T) {
		v := []int64{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]int64 nil", func(t *testing.T) {
		var v []int64 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int64 non-nil", func(t *testing.T) {
		s := []int64{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]int64 nil", func(t *testing.T) {
		var ps *[]int64 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int64", func(t *testing.T) {
		a, b, c := int64(1), int64(2), int64(3)
		v := []*int64{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int64 with nil", func(t *testing.T) {
		a, b := int64(1), int64(2)
		v := []*int64{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*int64 all nil", func(t *testing.T) {
		v := []*int64{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int64 non-nil", func(t *testing.T) {
		a, b := int64(1), int64(2)
		s := []*int64{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*int64 nil", func(t *testing.T) {
		var ps *[]*int64 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]int64", func(t *testing.T) {
		v := [3]int64{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int64 non-nil", func(t *testing.T) {
		v := [3]int64{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]int64 nil", func(t *testing.T) {
		var pv *[3]int64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int64", func(t *testing.T) {
		a, b, c := int64(1), int64(2), int64(3)
		v := [3]*int64{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int64 with nil", func(t *testing.T) {
		a, b := int64(1), int64(2)
		v := [3]*int64{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*int64 all nil", func(t *testing.T) {
		v := [3]*int64{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int64 non-nil", func(t *testing.T) {
		a, b := int64(1), int64(2)
		v := [3]*int64{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*int64 nil", func(t *testing.T) {
		var pv *[3]*int64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]int64", func(t *testing.T) {
		var v [100]int64
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szInt64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]int64 nil", func(t *testing.T) {
		var pv *[100]int64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*int64 all nil", func(t *testing.T) {
		var v [100]*int64
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*int64 nil", func(t *testing.T) {
		var pv *[100]*int64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_uint(t *testing.T) {
	t.Run("uint", func(t *testing.T) {
		v := uint(42)
		got := reflect.EstimatePayloadOf(v)
		want := szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint non-nil", func(t *testing.T) {
		v := uint(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint nil", func(t *testing.T) {
		var pv *uint = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint", func(t *testing.T) {
		v := []uint{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint empty", func(t *testing.T) {
		v := []uint{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint nil", func(t *testing.T) {
		var v []uint = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint non-nil", func(t *testing.T) {
		s := []uint{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint nil", func(t *testing.T) {
		var ps *[]uint = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint", func(t *testing.T) {
		a, b, c := uint(1), uint(2), uint(3)
		v := []*uint{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint with nil", func(t *testing.T) {
		a, b := uint(1), uint(2)
		v := []*uint{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint all nil", func(t *testing.T) {
		v := []*uint{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint non-nil", func(t *testing.T) {
		a, b := uint(1), uint(2)
		s := []*uint{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint nil", func(t *testing.T) {
		var ps *[]*uint = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]uint", func(t *testing.T) {
		v := [3]uint{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint non-nil", func(t *testing.T) {
		v := [3]uint{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint nil", func(t *testing.T) {
		var pv *[3]uint = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint", func(t *testing.T) {
		a, b, c := uint(1), uint(2), uint(3)
		v := [3]*uint{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint with nil", func(t *testing.T) {
		a, b := uint(1), uint(2)
		v := [3]*uint{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint all nil", func(t *testing.T) {
		v := [3]*uint{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint non-nil", func(t *testing.T) {
		a, b := uint(1), uint(2)
		v := [3]*uint{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint nil", func(t *testing.T) {
		var pv *[3]*uint = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]uint", func(t *testing.T) {
		var v [100]uint
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szUint
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]uint nil", func(t *testing.T) {
		var pv *[100]uint = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*uint all nil", func(t *testing.T) {
		var v [100]*uint
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*uint nil", func(t *testing.T) {
		var pv *[100]*uint = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_uint8(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		v := uint8(42)
		got := reflect.EstimatePayloadOf(v)
		want := szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint8 non-nil", func(t *testing.T) {
		v := uint8(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint8 nil", func(t *testing.T) {
		var pv *uint8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint8", func(t *testing.T) {
		v := []uint8{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint8 empty", func(t *testing.T) {
		v := []uint8{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint8 nil", func(t *testing.T) {
		var v []uint8 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint8 non-nil", func(t *testing.T) {
		s := []uint8{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint8 nil", func(t *testing.T) {
		var ps *[]uint8 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint8", func(t *testing.T) {
		a, b, c := uint8(1), uint8(2), uint8(3)
		v := []*uint8{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint8 with nil", func(t *testing.T) {
		a, b := uint8(1), uint8(2)
		v := []*uint8{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint8 all nil", func(t *testing.T) {
		v := []*uint8{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint8 non-nil", func(t *testing.T) {
		a, b := uint8(1), uint8(2)
		s := []*uint8{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint8 nil", func(t *testing.T) {
		var ps *[]*uint8 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]uint8", func(t *testing.T) {
		v := [3]uint8{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint8 non-nil", func(t *testing.T) {
		v := [3]uint8{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint8 nil", func(t *testing.T) {
		var pv *[3]uint8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint8", func(t *testing.T) {
		a, b, c := uint8(1), uint8(2), uint8(3)
		v := [3]*uint8{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint8 with nil", func(t *testing.T) {
		a, b := uint8(1), uint8(2)
		v := [3]*uint8{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint8 all nil", func(t *testing.T) {
		v := [3]*uint8{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint8 non-nil", func(t *testing.T) {
		a, b := uint8(1), uint8(2)
		v := [3]*uint8{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint8 nil", func(t *testing.T) {
		var pv *[3]*uint8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]uint8", func(t *testing.T) {
		var v [100]uint8
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szUint8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]uint8 nil", func(t *testing.T) {
		var pv *[100]uint8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*uint8 all nil", func(t *testing.T) {
		var v [100]*uint8
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*uint8 nil", func(t *testing.T) {
		var pv *[100]*uint8 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_uint16(t *testing.T) {
	t.Run("uint16", func(t *testing.T) {
		v := uint16(42)
		got := reflect.EstimatePayloadOf(v)
		want := szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint16 non-nil", func(t *testing.T) {
		v := uint16(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint16 nil", func(t *testing.T) {
		var pv *uint16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint16", func(t *testing.T) {
		v := []uint16{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint16 empty", func(t *testing.T) {
		v := []uint16{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint16 nil", func(t *testing.T) {
		var v []uint16 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint16 non-nil", func(t *testing.T) {
		s := []uint16{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint16 nil", func(t *testing.T) {
		var ps *[]uint16 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint16", func(t *testing.T) {
		a, b, c := uint16(1), uint16(2), uint16(3)
		v := []*uint16{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint16 with nil", func(t *testing.T) {
		a, b := uint16(1), uint16(2)
		v := []*uint16{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint16 all nil", func(t *testing.T) {
		v := []*uint16{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint16 non-nil", func(t *testing.T) {
		a, b := uint16(1), uint16(2)
		s := []*uint16{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint16 nil", func(t *testing.T) {
		var ps *[]*uint16 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]uint16", func(t *testing.T) {
		v := [3]uint16{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint16 non-nil", func(t *testing.T) {
		v := [3]uint16{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint16 nil", func(t *testing.T) {
		var pv *[3]uint16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint16", func(t *testing.T) {
		a, b, c := uint16(1), uint16(2), uint16(3)
		v := [3]*uint16{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint16 with nil", func(t *testing.T) {
		a, b := uint16(1), uint16(2)
		v := [3]*uint16{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint16 all nil", func(t *testing.T) {
		v := [3]*uint16{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint16 non-nil", func(t *testing.T) {
		a, b := uint16(1), uint16(2)
		v := [3]*uint16{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint16 nil", func(t *testing.T) {
		var pv *[3]*uint16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]uint16", func(t *testing.T) {
		var v [100]uint16
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szUint16
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]uint16 nil", func(t *testing.T) {
		var pv *[100]uint16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*uint16 all nil", func(t *testing.T) {
		var v [100]*uint16
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*uint16 nil", func(t *testing.T) {
		var pv *[100]*uint16 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_uint32(t *testing.T) {
	t.Run("uint32", func(t *testing.T) {
		v := uint32(42)
		got := reflect.EstimatePayloadOf(v)
		want := szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint32 non-nil", func(t *testing.T) {
		v := uint32(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint32 nil", func(t *testing.T) {
		var pv *uint32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint32", func(t *testing.T) {
		v := []uint32{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint32 empty", func(t *testing.T) {
		v := []uint32{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint32 nil", func(t *testing.T) {
		var v []uint32 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint32 non-nil", func(t *testing.T) {
		s := []uint32{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint32 nil", func(t *testing.T) {
		var ps *[]uint32 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint32", func(t *testing.T) {
		a, b, c := uint32(1), uint32(2), uint32(3)
		v := []*uint32{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint32 with nil", func(t *testing.T) {
		a, b := uint32(1), uint32(2)
		v := []*uint32{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint32 all nil", func(t *testing.T) {
		v := []*uint32{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint32 non-nil", func(t *testing.T) {
		a, b := uint32(1), uint32(2)
		s := []*uint32{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint32 nil", func(t *testing.T) {
		var ps *[]*uint32 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]uint32", func(t *testing.T) {
		v := [3]uint32{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint32 non-nil", func(t *testing.T) {
		v := [3]uint32{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint32 nil", func(t *testing.T) {
		var pv *[3]uint32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint32", func(t *testing.T) {
		a, b, c := uint32(1), uint32(2), uint32(3)
		v := [3]*uint32{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint32 with nil", func(t *testing.T) {
		a, b := uint32(1), uint32(2)
		v := [3]*uint32{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint32 all nil", func(t *testing.T) {
		v := [3]*uint32{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint32 non-nil", func(t *testing.T) {
		a, b := uint32(1), uint32(2)
		v := [3]*uint32{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint32 nil", func(t *testing.T) {
		var pv *[3]*uint32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]uint32", func(t *testing.T) {
		var v [100]uint32
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szUint32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]uint32 nil", func(t *testing.T) {
		var pv *[100]uint32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*uint32 all nil", func(t *testing.T) {
		var v [100]*uint32
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*uint32 nil", func(t *testing.T) {
		var pv *[100]*uint32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_uint64(t *testing.T) {
	t.Run("uint64", func(t *testing.T) {
		v := uint64(42)
		got := reflect.EstimatePayloadOf(v)
		want := szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint64 non-nil", func(t *testing.T) {
		v := uint64(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uint64 nil", func(t *testing.T) {
		var pv *uint64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint64", func(t *testing.T) {
		v := []uint64{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint64 empty", func(t *testing.T) {
		v := []uint64{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uint64 nil", func(t *testing.T) {
		var v []uint64 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint64 non-nil", func(t *testing.T) {
		s := []uint64{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uint64 nil", func(t *testing.T) {
		var ps *[]uint64 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint64", func(t *testing.T) {
		a, b, c := uint64(1), uint64(2), uint64(3)
		v := []*uint64{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint64 with nil", func(t *testing.T) {
		a, b := uint64(1), uint64(2)
		v := []*uint64{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uint64 all nil", func(t *testing.T) {
		v := []*uint64{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint64 non-nil", func(t *testing.T) {
		a, b := uint64(1), uint64(2)
		s := []*uint64{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uint64 nil", func(t *testing.T) {
		var ps *[]*uint64 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]uint64", func(t *testing.T) {
		v := [3]uint64{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint64 non-nil", func(t *testing.T) {
		v := [3]uint64{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uint64 nil", func(t *testing.T) {
		var pv *[3]uint64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint64", func(t *testing.T) {
		a, b, c := uint64(1), uint64(2), uint64(3)
		v := [3]*uint64{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint64 with nil", func(t *testing.T) {
		a, b := uint64(1), uint64(2)
		v := [3]*uint64{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uint64 all nil", func(t *testing.T) {
		v := [3]*uint64{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint64 non-nil", func(t *testing.T) {
		a, b := uint64(1), uint64(2)
		v := [3]*uint64{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uint64 nil", func(t *testing.T) {
		var pv *[3]*uint64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]uint64", func(t *testing.T) {
		var v [100]uint64
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szUint64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]uint64 nil", func(t *testing.T) {
		var pv *[100]uint64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*uint64 all nil", func(t *testing.T) {
		var v [100]*uint64
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*uint64 nil", func(t *testing.T) {
		var pv *[100]*uint64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_uintptr(t *testing.T) {
	t.Run("uintptr", func(t *testing.T) {
		v := uintptr(42)
		got := reflect.EstimatePayloadOf(v)
		want := szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uintptr non-nil", func(t *testing.T) {
		v := uintptr(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*uintptr nil", func(t *testing.T) {
		var pv *uintptr = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uintptr", func(t *testing.T) {
		v := []uintptr{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uintptr empty", func(t *testing.T) {
		v := []uintptr{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]uintptr nil", func(t *testing.T) {
		var v []uintptr = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uintptr non-nil", func(t *testing.T) {
		s := []uintptr{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]uintptr nil", func(t *testing.T) {
		var ps *[]uintptr = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uintptr", func(t *testing.T) {
		a, b, c := uintptr(1), uintptr(2), uintptr(3)
		v := []*uintptr{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uintptr with nil", func(t *testing.T) {
		a, b := uintptr(1), uintptr(2)
		v := []*uintptr{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*uintptr all nil", func(t *testing.T) {
		v := []*uintptr{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uintptr non-nil", func(t *testing.T) {
		a, b := uintptr(1), uintptr(2)
		s := []*uintptr{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*uintptr nil", func(t *testing.T) {
		var ps *[]*uintptr = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]uintptr", func(t *testing.T) {
		v := [3]uintptr{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uintptr non-nil", func(t *testing.T) {
		v := [3]uintptr{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]uintptr nil", func(t *testing.T) {
		var pv *[3]uintptr = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uintptr", func(t *testing.T) {
		a, b, c := uintptr(1), uintptr(2), uintptr(3)
		v := [3]*uintptr{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uintptr with nil", func(t *testing.T) {
		a, b := uintptr(1), uintptr(2)
		v := [3]*uintptr{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*uintptr all nil", func(t *testing.T) {
		v := [3]*uintptr{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uintptr non-nil", func(t *testing.T) {
		a, b := uintptr(1), uintptr(2)
		v := [3]*uintptr{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*uintptr nil", func(t *testing.T) {
		var pv *[3]*uintptr = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]uintptr", func(t *testing.T) {
		var v [100]uintptr
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szUintptr
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]uintptr nil", func(t *testing.T) {
		var pv *[100]uintptr = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*uintptr all nil", func(t *testing.T) {
		var v [100]*uintptr
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*uintptr nil", func(t *testing.T) {
		var pv *[100]*uintptr = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_float32(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		v := float32(42.0)
		got := reflect.EstimatePayloadOf(v)
		want := szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*float32 non-nil", func(t *testing.T) {
		v := float32(42.0)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*float32 nil", func(t *testing.T) {
		var pv *float32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]float32", func(t *testing.T) {
		v := []float32{1.0, 2.0, 3.0}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]float32 empty", func(t *testing.T) {
		v := []float32{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]float32 nil", func(t *testing.T) {
		var v []float32 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]float32 non-nil", func(t *testing.T) {
		s := []float32{1.0, 2.0, 3.0}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]float32 nil", func(t *testing.T) {
		var ps *[]float32 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*float32", func(t *testing.T) {
		a, b, c := float32(1.0), float32(2.0), float32(3.0)
		v := []*float32{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*float32 with nil", func(t *testing.T) {
		a, b := float32(1.0), float32(2.0)
		v := []*float32{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*float32 all nil", func(t *testing.T) {
		v := []*float32{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*float32 non-nil", func(t *testing.T) {
		a, b := float32(1.0), float32(2.0)
		s := []*float32{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*float32 nil", func(t *testing.T) {
		var ps *[]*float32 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]float32", func(t *testing.T) {
		v := [3]float32{1.0, 2.0, 3.0}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]float32 non-nil", func(t *testing.T) {
		v := [3]float32{1.0, 2.0, 3.0}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]float32 nil", func(t *testing.T) {
		var pv *[3]float32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*float32", func(t *testing.T) {
		a, b, c := float32(1.0), float32(2.0), float32(3.0)
		v := [3]*float32{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*float32 with nil", func(t *testing.T) {
		a, b := float32(1.0), float32(2.0)
		v := [3]*float32{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*float32 all nil", func(t *testing.T) {
		v := [3]*float32{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*float32 non-nil", func(t *testing.T) {
		a, b := float32(1.0), float32(2.0)
		v := [3]*float32{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*float32 nil", func(t *testing.T) {
		var pv *[3]*float32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]float32", func(t *testing.T) {
		var v [100]float32
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szFloat32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]float32 nil", func(t *testing.T) {
		var pv *[100]float32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*float32 all nil", func(t *testing.T) {
		var v [100]*float32
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*float32 nil", func(t *testing.T) {
		var pv *[100]*float32 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_float64(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		v := float64(42.0)
		got := reflect.EstimatePayloadOf(v)
		want := szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*float64 non-nil", func(t *testing.T) {
		v := float64(42.0)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*float64 nil", func(t *testing.T) {
		var pv *float64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]float64", func(t *testing.T) {
		v := []float64{1.0, 2.0, 3.0}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]float64 empty", func(t *testing.T) {
		v := []float64{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]float64 nil", func(t *testing.T) {
		var v []float64 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]float64 non-nil", func(t *testing.T) {
		s := []float64{1.0, 2.0, 3.0}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]float64 nil", func(t *testing.T) {
		var ps *[]float64 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*float64", func(t *testing.T) {
		a, b, c := float64(1.0), float64(2.0), float64(3.0)
		v := []*float64{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*float64 with nil", func(t *testing.T) {
		a, b := float64(1.0), float64(2.0)
		v := []*float64{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*float64 all nil", func(t *testing.T) {
		v := []*float64{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*float64 non-nil", func(t *testing.T) {
		a, b := float64(1.0), float64(2.0)
		s := []*float64{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*float64 nil", func(t *testing.T) {
		var ps *[]*float64 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]float64", func(t *testing.T) {
		v := [3]float64{1.0, 2.0, 3.0}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]float64 non-nil", func(t *testing.T) {
		v := [3]float64{1.0, 2.0, 3.0}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]float64 nil", func(t *testing.T) {
		var pv *[3]float64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*float64", func(t *testing.T) {
		a, b, c := float64(1.0), float64(2.0), float64(3.0)
		v := [3]*float64{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*float64 with nil", func(t *testing.T) {
		a, b := float64(1.0), float64(2.0)
		v := [3]*float64{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*float64 all nil", func(t *testing.T) {
		v := [3]*float64{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*float64 non-nil", func(t *testing.T) {
		a, b := float64(1.0), float64(2.0)
		v := [3]*float64{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*float64 nil", func(t *testing.T) {
		var pv *[3]*float64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]float64", func(t *testing.T) {
		var v [100]float64
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szFloat64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]float64 nil", func(t *testing.T) {
		var pv *[100]float64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*float64 all nil", func(t *testing.T) {
		var v [100]*float64
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*float64 nil", func(t *testing.T) {
		var pv *[100]*float64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_complex64(t *testing.T) {
	t.Run("complex64", func(t *testing.T) {
		v := complex64(42)
		got := reflect.EstimatePayloadOf(v)
		want := szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*complex64 non-nil", func(t *testing.T) {
		v := complex64(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*complex64 nil", func(t *testing.T) {
		var pv *complex64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]complex64", func(t *testing.T) {
		v := []complex64{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]complex64 empty", func(t *testing.T) {
		v := []complex64{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]complex64 nil", func(t *testing.T) {
		var v []complex64 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]complex64 non-nil", func(t *testing.T) {
		s := []complex64{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]complex64 nil", func(t *testing.T) {
		var ps *[]complex64 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*complex64", func(t *testing.T) {
		a, b, c := complex64(1), complex64(2), complex64(3)
		v := []*complex64{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*complex64 with nil", func(t *testing.T) {
		a, b := complex64(1), complex64(2)
		v := []*complex64{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*complex64 all nil", func(t *testing.T) {
		v := []*complex64{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*complex64 non-nil", func(t *testing.T) {
		a, b := complex64(1), complex64(2)
		s := []*complex64{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*complex64 nil", func(t *testing.T) {
		var ps *[]*complex64 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]complex64", func(t *testing.T) {
		v := [3]complex64{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]complex64 non-nil", func(t *testing.T) {
		v := [3]complex64{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]complex64 nil", func(t *testing.T) {
		var pv *[3]complex64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*complex64", func(t *testing.T) {
		a, b, c := complex64(1), complex64(2), complex64(3)
		v := [3]*complex64{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*complex64 with nil", func(t *testing.T) {
		a, b := complex64(1), complex64(2)
		v := [3]*complex64{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*complex64 all nil", func(t *testing.T) {
		v := [3]*complex64{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*complex64 non-nil", func(t *testing.T) {
		a, b := complex64(1), complex64(2)
		v := [3]*complex64{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*complex64 nil", func(t *testing.T) {
		var pv *[3]*complex64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]complex64", func(t *testing.T) {
		var v [100]complex64
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szComplex64
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]complex64 nil", func(t *testing.T) {
		var pv *[100]complex64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*complex64 all nil", func(t *testing.T) {
		var v [100]*complex64
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*complex64 nil", func(t *testing.T) {
		var pv *[100]*complex64 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_complex128(t *testing.T) {
	t.Run("complex128", func(t *testing.T) {
		v := complex128(42)
		got := reflect.EstimatePayloadOf(v)
		want := szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*complex128 non-nil", func(t *testing.T) {
		v := complex128(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*complex128 nil", func(t *testing.T) {
		var pv *complex128 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]complex128", func(t *testing.T) {
		v := []complex128{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]complex128 empty", func(t *testing.T) {
		v := []complex128{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]complex128 nil", func(t *testing.T) {
		var v []complex128 = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]complex128 non-nil", func(t *testing.T) {
		s := []complex128{1, 2, 3}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 3 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]complex128 nil", func(t *testing.T) {
		var ps *[]complex128 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*complex128", func(t *testing.T) {
		a, b, c := complex128(1), complex128(2), complex128(3)
		v := []*complex128{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*complex128 with nil", func(t *testing.T) {
		a, b := complex128(1), complex128(2)
		v := []*complex128{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*complex128 all nil", func(t *testing.T) {
		v := []*complex128{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*complex128 non-nil", func(t *testing.T) {
		a, b := complex128(1), complex128(2)
		s := []*complex128{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*complex128 nil", func(t *testing.T) {
		var ps *[]*complex128 = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]complex128", func(t *testing.T) {
		v := [3]complex128{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]complex128 non-nil", func(t *testing.T) {
		v := [3]complex128{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]complex128 nil", func(t *testing.T) {
		var pv *[3]complex128 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*complex128", func(t *testing.T) {
		a, b, c := complex128(1), complex128(2), complex128(3)
		v := [3]*complex128{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*complex128 with nil", func(t *testing.T) {
		a, b := complex128(1), complex128(2)
		v := [3]*complex128{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*complex128 all nil", func(t *testing.T) {
		v := [3]*complex128{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*complex128 non-nil", func(t *testing.T) {
		a, b := complex128(1), complex128(2)
		v := [3]*complex128{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*complex128 nil", func(t *testing.T) {
		var pv *[3]*complex128 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]complex128", func(t *testing.T) {
		var v [100]complex128
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szComplex128
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]complex128 nil", func(t *testing.T) {
		var pv *[100]complex128 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*complex128 all nil", func(t *testing.T) {
		var v [100]*complex128
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*complex128 nil", func(t *testing.T) {
		var pv *[100]*complex128 = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_bool(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		v := true
		got := reflect.EstimatePayloadOf(v)
		want := szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*bool non-nil", func(t *testing.T) {
		v := true
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*bool nil", func(t *testing.T) {
		var pv *bool = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]bool", func(t *testing.T) {
		v := []bool{true, false, true}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]bool empty", func(t *testing.T) {
		v := []bool{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]bool nil", func(t *testing.T) {
		var v []bool = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]bool non-nil", func(t *testing.T) {
		s := []bool{true, false}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]bool nil", func(t *testing.T) {
		var ps *[]bool = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*bool", func(t *testing.T) {
		a, b, c := true, false, true
		v := []*bool{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*bool with nil", func(t *testing.T) {
		a, b := true, false
		v := []*bool{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*bool all nil", func(t *testing.T) {
		v := []*bool{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*bool non-nil", func(t *testing.T) {
		a, b := true, false
		s := []*bool{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*bool nil", func(t *testing.T) {
		var ps *[]*bool = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]bool", func(t *testing.T) {
		v := [3]bool{true, false, true}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]bool non-nil", func(t *testing.T) {
		v := [3]bool{true, false, true}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]bool nil", func(t *testing.T) {
		var pv *[3]bool = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*bool", func(t *testing.T) {
		a, b, c := true, false, true
		v := [3]*bool{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*bool with nil", func(t *testing.T) {
		a, b := true, false
		v := [3]*bool{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*bool all nil", func(t *testing.T) {
		v := [3]*bool{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*bool non-nil", func(t *testing.T) {
		a, b := true, false
		v := [3]*bool{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*bool nil", func(t *testing.T) {
		var pv *[3]*bool = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]bool", func(t *testing.T) {
		var v [100]bool
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szBool
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]bool nil", func(t *testing.T) {
		var pv *[100]bool = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*bool all nil", func(t *testing.T) {
		var v [100]*bool
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*bool nil", func(t *testing.T) {
		var pv *[100]*bool = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_string(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v := "hello"
		got := reflect.EstimatePayloadOf(v)
		want := 5
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("string empty", func(t *testing.T) {
		v := ""
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*string non-nil", func(t *testing.T) {
		v := "hello"
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 5
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*string nil", func(t *testing.T) {
		var pv *string = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]string", func(t *testing.T) {
		v := []string{"he", "llo", ""}
		got := reflect.EstimatePayloadOf(v)
		want := 2 + 3 + 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]string empty", func(t *testing.T) {
		v := []string{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]string nil", func(t *testing.T) {
		var v []string = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]string non-nil", func(t *testing.T) {
		s := []string{"he", "llo"}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 + 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]string nil", func(t *testing.T) {
		var ps *[]string = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*string", func(t *testing.T) {
		a, b, c := "he", "llo", ""
		v := []*string{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 2 + 3 + 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*string with nil", func(t *testing.T) {
		a, b := "he", "llo"
		v := []*string{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 + 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*string all nil", func(t *testing.T) {
		v := []*string{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*string non-nil", func(t *testing.T) {
		a, b := "he", "llo"
		s := []*string{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 + 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*string nil", func(t *testing.T) {
		var ps *[]*string = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]string", func(t *testing.T) {
		v := [3]string{"he", "llo", ""}
		got := reflect.EstimatePayloadOf(v)
		want := 2 + 3 + 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]string non-nil", func(t *testing.T) {
		v := [3]string{"he", "llo", ""}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 + 3 + 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]string nil", func(t *testing.T) {
		var pv *[3]string = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*string", func(t *testing.T) {
		a, b, c := "he", "llo", ""
		v := [3]*string{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 2 + 3 + 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*string with nil", func(t *testing.T) {
		a, b := "he", "llo"
		v := [3]*string{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 + 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*string all nil", func(t *testing.T) {
		v := [3]*string{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*string non-nil", func(t *testing.T) {
		a, b := "he", "llo"
		v := [3]*string{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 + 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*string nil", func(t *testing.T) {
		var pv *[3]*string = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]string", func(t *testing.T) {
		var v [100]string // all empty
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]string nil", func(t *testing.T) {
		var pv *[100]string = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*string all nil", func(t *testing.T) {
		var v [100]*string
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*string nil", func(t *testing.T) {
		var pv *[100]*string = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_timeTime(t *testing.T) {
	t.Run("time.Time", func(t *testing.T) {
		v := time.Now()
		got := reflect.EstimatePayloadOf(v)
		want := szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*time.Time non-nil", func(t *testing.T) {
		v := time.Now()
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*time.Time nil", func(t *testing.T) {
		var pv *time.Time = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]time.Time", func(t *testing.T) {
		v := []time.Time{time.Now(), time.Now(), time.Now()}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]time.Time empty", func(t *testing.T) {
		v := []time.Time{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]time.Time nil", func(t *testing.T) {
		var v []time.Time = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]time.Time non-nil", func(t *testing.T) {
		s := []time.Time{time.Now(), time.Now()}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]time.Time nil", func(t *testing.T) {
		var ps *[]time.Time = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*time.Time", func(t *testing.T) {
		a := time.Now()
		b := time.Now()
		c := time.Now()
		v := []*time.Time{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*time.Time with nil", func(t *testing.T) {
		a, b := time.Now(), time.Now()
		v := []*time.Time{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*time.Time all nil", func(t *testing.T) {
		v := []*time.Time{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*time.Time non-nil", func(t *testing.T) {
		a, b := time.Now(), time.Now()
		s := []*time.Time{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*time.Time nil", func(t *testing.T) {
		var ps *[]*time.Time = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]time.Time", func(t *testing.T) {
		v := [3]time.Time{time.Now(), time.Now(), time.Now()}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]time.Time non-nil", func(t *testing.T) {
		v := [3]time.Time{time.Now(), time.Now(), time.Now()}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]time.Time nil", func(t *testing.T) {
		var pv *[3]time.Time = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*time.Time", func(t *testing.T) {
		a, b, c := time.Now(), time.Now(), time.Now()
		v := [3]*time.Time{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*time.Time with nil", func(t *testing.T) {
		a, b := time.Now(), time.Now()
		v := [3]*time.Time{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*time.Time all nil", func(t *testing.T) {
		v := [3]*time.Time{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*time.Time non-nil", func(t *testing.T) {
		a, b := time.Now(), time.Now()
		v := [3]*time.Time{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*time.Time nil", func(t *testing.T) {
		var pv *[3]*time.Time = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]time.Time", func(t *testing.T) {
		var v [100]time.Time
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szTime
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]time.Time nil", func(t *testing.T) {
		var pv *[100]time.Time = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*time.Time all nil", func(t *testing.T) {
		var v [100]*time.Time
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*time.Time nil", func(t *testing.T) {
		var pv *[100]*time.Time = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_timeDuration(t *testing.T) {
	t.Run("time.Duration", func(t *testing.T) {
		v := time.Duration(42)
		got := reflect.EstimatePayloadOf(v)
		want := szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*time.Duration non-nil", func(t *testing.T) {
		v := time.Duration(42)
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*time.Duration nil", func(t *testing.T) {
		var pv *time.Duration = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]time.Duration", func(t *testing.T) {
		v := []time.Duration{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]time.Duration empty", func(t *testing.T) {
		v := []time.Duration{}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]time.Duration nil", func(t *testing.T) {
		var v []time.Duration = nil
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]time.Duration non-nil", func(t *testing.T) {
		s := []time.Duration{1, 2}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]time.Duration nil", func(t *testing.T) {
		var ps *[]time.Duration = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*time.Duration", func(t *testing.T) {
		a, b, c := time.Duration(1), time.Duration(2), time.Duration(3)
		v := []*time.Duration{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*time.Duration with nil", func(t *testing.T) {
		a, b := time.Duration(1), time.Duration(2)
		v := []*time.Duration{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*time.Duration all nil", func(t *testing.T) {
		v := []*time.Duration{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*time.Duration non-nil", func(t *testing.T) {
		a, b := time.Duration(1), time.Duration(2)
		s := []*time.Duration{&a, nil, &b}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := 2 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*time.Duration nil", func(t *testing.T) {
		var ps *[]*time.Duration = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]time.Duration", func(t *testing.T) {
		v := [3]time.Duration{1, 2, 3}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]time.Duration non-nil", func(t *testing.T) {
		v := [3]time.Duration{1, 2, 3}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 3 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]time.Duration nil", func(t *testing.T) {
		var pv *[3]time.Duration = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*time.Duration", func(t *testing.T) {
		a, b, c := time.Duration(1), time.Duration(2), time.Duration(3)
		v := [3]*time.Duration{&a, &b, &c}
		got := reflect.EstimatePayloadOf(v)
		want := 3 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*time.Duration with nil", func(t *testing.T) {
		a, b := time.Duration(1), time.Duration(2)
		v := [3]*time.Duration{&a, nil, &b}
		got := reflect.EstimatePayloadOf(v)
		want := 2 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*time.Duration all nil", func(t *testing.T) {
		v := [3]*time.Duration{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*time.Duration non-nil", func(t *testing.T) {
		a, b := time.Duration(1), time.Duration(2)
		v := [3]*time.Duration{&a, nil, &b}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := 2 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*time.Duration nil", func(t *testing.T) {
		var pv *[3]*time.Duration = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]time.Duration", func(t *testing.T) {
		var v [100]time.Duration
		got := reflect.EstimatePayloadOf(v)
		want := 100 * szDuration
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]time.Duration nil", func(t *testing.T) {
		var pv *[100]time.Duration = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*time.Duration all nil", func(t *testing.T) {
		var v [100]*time.Duration
		got := reflect.EstimatePayloadOf(v)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*time.Duration nil", func(t *testing.T) {
		var pv *[100]*time.Duration = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestEstimatePayloadOf_CustomStruct(t *testing.T) {
	t.Run("CustomStruct", func(t *testing.T) {
		v := CustomStruct{}
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*CustomStruct non-nil", func(t *testing.T) {
		v := CustomStruct{}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*CustomStruct nil", func(t *testing.T) {
		var pv *CustomStruct = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]CustomStruct", func(t *testing.T) {
		v := []CustomStruct{{}, {}}
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]CustomStruct empty", func(t *testing.T) {
		v := []CustomStruct{}
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]CustomStruct nil", func(t *testing.T) {
		var v []CustomStruct = nil
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload // because we the EstimatePayloadOf don't know about CustomStruct
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]CustomStruct non-nil", func(t *testing.T) {
		s := []CustomStruct{{}}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]CustomStruct nil", func(t *testing.T) {
		var ps *[]CustomStruct = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*CustomStruct", func(t *testing.T) {
		a := CustomStruct{}
		v := []*CustomStruct{&a, nil}
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[]*CustomStruct all nil", func(t *testing.T) {
		v := []*CustomStruct{nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload // since fallback to payloadArray, and struct unknown, Err
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*CustomStruct non-nil", func(t *testing.T) {
		s := []*CustomStruct{nil}
		ps := &s
		got := reflect.EstimatePayloadOf(ps)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[]*CustomStruct nil", func(t *testing.T) {
		var ps *[]*CustomStruct = nil
		got := reflect.EstimatePayloadOf(ps)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]CustomStruct", func(t *testing.T) {
		v := [3]CustomStruct{{}, {}, {}}
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]CustomStruct non-nil", func(t *testing.T) {
		v := [3]CustomStruct{{}, {}, {}}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]CustomStruct nil", func(t *testing.T) {
		var pv *[3]CustomStruct = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*CustomStruct", func(t *testing.T) {
		a := CustomStruct{}
		v := [3]*CustomStruct{&a, nil, &a}
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[3]*CustomStruct all nil", func(t *testing.T) {
		v := [3]*CustomStruct{nil, nil, nil}
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload // Err for unknown struct
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*CustomStruct non-nil", func(t *testing.T) {
		v := [3]*CustomStruct{nil, nil, nil}
		pv := &v
		got := reflect.EstimatePayloadOf(pv)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[3]*CustomStruct nil", func(t *testing.T) {
		var pv *[3]*CustomStruct = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]CustomStruct", func(t *testing.T) {
		var v [100]CustomStruct
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]CustomStruct nil", func(t *testing.T) {
		var pv *[100]CustomStruct = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("[100]*CustomStruct all nil", func(t *testing.T) {
		var v [100]*CustomStruct
		got := reflect.EstimatePayloadOf(v)
		want := reflect.ErrFailEstimatePayload
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("*[100]*CustomStruct nil", func(t *testing.T) {
		var pv *[100]*CustomStruct = nil
		got := reflect.EstimatePayloadOf(pv)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
