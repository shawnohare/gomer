package gomer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorError(t *testing.T) {
	var err *Error
	assert.Equal(t, "", err.Error())

	err = &Error{Index: 0, Key: "key", Err: errors.New("test")}
	assert.NotEqual(t, "", err.Error())
}
