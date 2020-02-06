// +build go1.13

package errors

import "fmt"

//Annotate an error with annotation
func Annotate(err error, annotation string) error {
	if err != nil {
		return fmt.Errorf("%s %w", annotation, err)
	}
	return nil
}

//Annotatef an error with annotation
func Annotatef(err error, format string, args ...interface{}) error {
	return Annotate(err, fmt.Sprintf(format, args...))
}
