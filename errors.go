package gomer

// Errors represents a collection of errors over an iterable, such as slices
// or maps.
type Errors struct {
	errs []*Error
}

// Slice returns the underlying collection of errors.
func (e *Errors) Slice() []*Error {
	if e == nil {
		return nil
	}
	return e.errs
}

// Len reports how many errors have been recorded.
func (e *Errors) Len() int {
	if e == nil {
		return 0
	}
	return len(e.errs)
}

// Empty reports whether the instance contains at least one non-nil error.
func (e *Errors) Empty() bool {
	for _, err := range e.Slice() {
		if err != nil && err.Err != nil {
			return false
		}
	}
	return true
}

// Add a new Error with the given index, key, and underlying error.
// Uniquness of the index and keys is not checked
func (e *Errors) Add(index int, key string, err error) *Errors {
	if e == nil {
		e = New()
	}
	e.errs = append(e.errs, &Error{index, key, err})
	return e
}

// AddNonNil is the same as Add but skips nil error inputs.
func (e *Errors) AddNonNil(index int, key string, err error) *Errors {
	if err == nil {
		return e
	}
	return e.Add(index, key, err)
}

// Error is defined so that the Errors struct implements the error interface.
func (e *Errors) Error() string {
	if e == nil {
		return ""
	}

	var errStr string
	for _, err := range e.errs {
		if s := err.Error(); s != "" {
			errStr += s + "\n"
		}
	}

	return errStr
}

// New returns an Error instance that is ready for use.
func New() *Errors {
	return new(Errors)
}
