package gomer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testErr error = errors.New("test")

func TestErrorsAdd(t *testing.T) {
	var errs *Errors
	errs = errs.Add(0, "0", errors.New("test"))
	assert.Len(t, errs.errs, 1)

	if errs.Len() > 0 {
		actual := errs.errs[0]
		expected := &Error{0, "0", errors.New("test")}
		assert.Equal(t, expected, actual)
	}
}

func TestErrorsAddNonNil(t *testing.T) {
	errs := New()
	assert.Equal(t, 0, errs.AddNonNil(0, "", nil).Len())
	assert.Equal(t, 1, errs.AddNonNil(0, "", testErr).Len())
}

func TestErrorsEmpty(t *testing.T) {
	var errs *Errors
	assert.True(t, errs.Empty())
	assert.True(t, errs.Add(0, "", nil).Empty())
	assert.False(t, errs.Add(0, "", testErr).Empty())
}

func TestErrorsLenNil(t *testing.T) {
	var errs *Errors
	assert.Equal(t, 0, errs.Len())
}

func TestErrorsError(t *testing.T) {
	var errs *Errors
	assert.Equal(t, "", errs.Error())

	errs = errs.Add(0, "", nil)
	assert.Equal(t, "", errs.Error())

	errs = errs.Add(0, "", testErr)
	assert.NotEqual(t, "", errs.Error())
}
