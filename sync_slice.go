package multierror

import (
	"github.com/xaionaro-go/synctools"
)

// SyncSlice is a synchronous variant of Slice. It is safe
// to its call methods from multiple goroutines without external locks.
type SyncSlice struct {
	Slice

	locker synctools.RWMutex
}

// Error implements interface `error`.
func (errs *SyncSlice) Error() string {
	var result string
	errs.locker.RLockDo(func() {
		result = errs.Slice.Error()
	})
	return result
}

// Unwrap implements the interface used by `errors.Unwrap()`.
//
// If the slice of errors contains only one error then this
// error is returned, otherwise `nil` is returned.
func (errs *SyncSlice) Unwrap() error {
	var result error
	errs.locker.RLockDo(func() {
		result = errs.Slice.Unwrap()
	})
	return result
}

// As implements the interface used by `errors.As()`.
//
// Calls `errors.As()` for each error of the slice until
// `errors.As()` will return `true`. If no `true` was returned
// by `errors.As()` then false is returned.
func (errs *SyncSlice) As(dst interface{}) bool {
	var result bool
	errs.locker.RLockDo(func() {
		result = errs.Slice.As(dst)
	})
	return result
}

// Is implements the interface used by `errors.Is()`.
//
// Calls `errors.Is()` for each error of the slice until
// `errors.Is()` will return `true`. If no `true` was returned
// by `errors.Is()` then false is returned.
func (errs *SyncSlice) Is(cmp error) bool {
	var result bool
	errs.locker.RLockDo(func() {
		result = errs.Slice.Is(cmp)
	})
	return result
}

// Add adds non-nil errors `addErrs` to the slice.
func (errs *SyncSlice) Add(addErrs ...error) {
	errs.locker.LockDo(func() {
		errs.Slice.Add(addErrs...)
	})
}

// ReturnValue returns untyped nil if the slice is empty
// and returns the slice if it is not empty.
//
// It is supposed to be used in `return`-s, like:
//
//    return err.ReturnValue()
func (errs *SyncSlice) ReturnValue() error {
	var result error
	errs.locker.RLockDo(func() {
		result = errs.Slice.ReturnValue()
	})
	return result
}
