package chanx_test

import (
	"reflect"
	"testing"

	"github.com/lif0/pkg/concurrency/chanx"
)

// TestToRecvChans verifies that ToRecvChans converts a slice of bidirectional channels to receive-only.
func TestToRecvChans(t *testing.T) {
	tests := []struct {
		name string
		in   []chan int
		want int // Expected length
	}{
		{
			name: "Empty slice",
			in:   []chan int{},
			want: 0,
		},
		{
			name: "Single channel",
			in:   []chan int{make(chan int)},
			want: 1,
		},
		{
			name: "Multiple channels",
			in:   []chan int{make(chan int), make(chan int), make(chan int)},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chanx.ToRecvChans(tt.in)
			if len(got) != tt.want {
				t.Errorf("ToRecvChans() len = %v, want %v", len(got), tt.want)
			}

			for i := range got {
				if got[i] != tt.in[i] {
					t.Errorf("ToRecvChans() element %d does not match original", i)
				}
			}

			expectedType := reflect.TypeOf([]<-chan int(nil))
			if reflect.TypeOf(got) != expectedType {
				t.Errorf("ToRecvChans() type = %v, want %v", reflect.TypeOf(got), expectedType)
			}
		})
	}
}

// TestToSendChans verifies that ToSendChans converts a slice of bidirectional channels to send-only.
func TestToSendChans(t *testing.T) {
	tests := []struct {
		name string
		in   []chan string
		want int // Expected length
	}{
		{
			name: "Empty slice",
			in:   []chan string{},
			want: 0,
		},
		{
			name: "Single channel",
			in:   []chan string{make(chan string)},
			want: 1,
		},
		{
			name: "Multiple channels",
			in:   []chan string{make(chan string), make(chan string)},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chanx.ToSendChans(tt.in)
			if len(got) != tt.want {
				t.Errorf("ToSendChans() len = %v, want %v", len(got), tt.want)
			}

			for i := range got {
				if got[i] != tt.in[i] {
					t.Errorf("ToSendChans() element %d does not match original", i)
				}
			}

			expectedType := reflect.TypeOf([]chan<- string(nil))
			if reflect.TypeOf(got) != expectedType {
				t.Errorf("ToSendChans() type = %v, want %v", reflect.TypeOf(got), expectedType)
			}
		})
	}
}
