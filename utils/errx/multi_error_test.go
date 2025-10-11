package errx_test

import (
	"fmt"
	"testing"

	"github.com/lif0/pkg/utils/errx"
)

func TestMultiError_Error(t *testing.T) {
	tests := []struct {
		name string
		errs errx.MultiError
		want string
	}{
		{
			name: "empty",
			errs: errx.MultiError{},
			want: "",
		},
		{
			name: "single error",
			errs: errx.MultiError{fmt.Errorf("error one")},
			want: "1 error(s) occurred:\n* error one",
		},
		{
			name: "multiple errors",
			errs: errx.MultiError{fmt.Errorf("error one"), fmt.Errorf("error two"), fmt.Errorf("error three")},
			want: "3 error(s) occurred:\n* error one\n* error two\n* error three",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.errs.Error(); got != tt.want {
				t.Errorf("MultiError.Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMultiError_Append(t *testing.T) {
	var me errx.MultiError

	// Append nil, should not add
	me.Append(nil)
	if len(me) != 0 {
		t.Errorf("After appending nil, len(me) = %d, want 0", len(me))
	}

	// Append first error
	err1 := fmt.Errorf("error one")
	me.Append(err1)
	if len(me) != 1 {
		t.Errorf("After appending first error, len(me) = %d, want 1", len(me))
	}
	if me[0] != err1 {
		t.Errorf("me[0] = %v, want %v", me[0], err1)
	}

	// Append nil again, should not add
	me.Append(nil)
	if len(me) != 1 {
		t.Errorf("After appending nil again, len(me) = %d, want 1", len(me))
	}

	// Append second error
	err2 := fmt.Errorf("error two")
	me.Append(err2)
	if len(me) != 2 {
		t.Errorf("After appending second error, len(me) = %d, want 2", len(me))
	}
	if me[1] != err2 {
		t.Errorf("me[1] = %v, want %v", me[1], err2)
	}
}

func TestMultiError_MaybeUnwrap(t *testing.T) {
	tests := []struct {
		name string
		errs errx.MultiError
		want error
	}{
		{
			name: "empty",
			errs: errx.MultiError{},
			want: nil,
		},
		{
			name: "single error",
			errs: errx.MultiError{fmt.Errorf("error one")},
			want: fmt.Errorf("error one"),
		},
		{
			name: "multiple errors",
			errs: errx.MultiError{fmt.Errorf("error one"), fmt.Errorf("error two")},
			want: errx.MultiError{fmt.Errorf("error one"), fmt.Errorf("error two")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.errs.MaybeUnwrap()
			switch tt.want.(type) {
			case errx.MultiError:
				if me, ok := got.(errx.MultiError); !ok || len(me) != len(tt.want.(errx.MultiError)) {
					t.Errorf("MultiError.MaybeUnwrap() = %v, want %v", got, tt.want)
				}
			default:
				if got != tt.want && (got == nil || tt.want == nil || got.Error() != tt.want.Error()) {
					t.Errorf("MultiError.MaybeUnwrap() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
