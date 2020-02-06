package errors

import (
	"errors"
	"fmt"

	stackerr "github.com/pkg/errors"
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

//As standard func
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

//Is standard func
func Is(err, target error) bool {
	return errors.Is(err, target)
}

//Unwrap standard func
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

//WithStack error
func WithStack(err error) error {
	return stackerr.WithStack(err)
}
