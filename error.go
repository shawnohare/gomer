package gomer

import "fmt"

// Error is a small wrapper around an error that records optional location
// data defined by the Index and Key fields.  The exact semantics of these
// fields are left to the user.
type Error struct {
	Index int
	Key   string
	Err   error
}

// Error is defined so that the Error struct implements the error interface.
func (err *Error) Error() string {
	if err == nil || err.Err == nil {
		return ""
	}

	return fmt.Sprintf(
		"Index: %d, Key: %s, Error: %s",
		err.Index,
		err.Key,
		err.Err.Error(),
	)
}
