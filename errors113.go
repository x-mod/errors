// +build go1.13

package errors

import (
	"errors"
)

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
