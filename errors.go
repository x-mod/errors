package errors

import (
	"errors"
	"fmt"
)

type wrapping interface {
	Unwrap() error
}

type coder interface {
	Value() int32
}

//New errorstring error
func New(err string) error {
	return errors.New(err)
}

//Errorf standard func
func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}
