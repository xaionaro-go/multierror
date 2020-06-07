package multierror

import (
	"errors"
	"fmt"
	"strings"
)

// Slice is a slice of errors which implements interface `error`.
type Slice []error

// Error implements interface `error`.
func (errs Slice) Error() string {
	var errDescriptions []string
	for _, err := range errs {
		errDescriptions = append(errDescriptions, err.Error())
	}

	return fmt.Sprintf("multiple errors:\n\t* %s\n", strings.Join(errDescriptions, "\n\t* "))
}

// Unwrap implements the interface used by `errors.Unwrap()`.
//
// If the slice of errors contains only one error then this
// error is returned, otherwise `nil` is returned.
func (errs Slice) Unwrap() error {
	if len(errs) == 1 {
		return errs[0]
	}

	return nil
}

// As implements the interface used by `errors.As()`.
//
// Calls `errors.As()` for each error of the slice until
// `errors.As()` will return `true`. If no `true` was returned
// by `errors.As()` then false is returned.
func (errs Slice) As(dst interface{}) bool {
	for _, err := range errs {
		if errors.As(err, dst) {
			return true
		}
	}

	return false
}

// Is implements the interface used by `errors.Is()`.
//
// Calls `errors.Is()` for each error of the slice until
// `errors.Is()` will return `true`. If no `true` was returned
// by `errors.Is()` then false is returned.
func (errs Slice) Is(cmp error) bool {
	for _, err := range errs {
		if errors.Is(err, cmp) {
			return true
		}
	}

	return false
}

// Add adds non-nil errors `addErrs` to the slice.
func (errs *Slice) Add(addErrs ...error) {
	for _, addErr := range addErrs {
		if addErr == nil {
			continue
		}

		*errs = append(*errs, addErr)
	}
}

// ReturnValue returns untyped nil if the slice is empty
// and returns the slice if it is not empty.
//
// It is supposed to be used in `return`-s, like:
//
//    return err.ReturnValue()
func (errs Slice) ReturnValue() error {
	if len(errs) == 0 {
		return nil
	}

	return errs
}
