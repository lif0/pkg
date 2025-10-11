package errx

import (
	"bytes"
	"fmt"
)

// MultiError is a slice of errors implementing the error interface.
type MultiError []error

// Error formats the contained errors as a bullet point list, preceded by the
// total number of errors. Note that this results in a multi-line string.
func (errs MultiError) Error() string {
	if len(errs) == 0 {
		return ""
	}
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%d error(s) occurred:", len(errs))
	for _, err := range errs {
		fmt.Fprintf(buf, "\n* %s", err)
	}
	return buf.String()
}

// Append appends the provided error if it is not nil.
func (errs *MultiError) Append(err error) {
	if err != nil {
		*errs = append(*errs, err)
	}
}

// MaybeUnwrap returns nil if len(errs) is 0. It returns the first and only
// contained error as error if len(errs is 1). In all other cases, it returns
// the MultiError directly. This is helpful for returning a MultiError in a way
// that only uses the MultiError if needed.
func (errs MultiError) MaybeUnwrap() error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errs
	}
}
