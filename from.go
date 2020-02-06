package errors

import (
	"errors"

	"google.golang.org/grpc/status"
)

//ValueFrom get code from the error
//support code value from grpc status
func ValueFrom(err error) int32 {
	if err != nil {
		for err != nil {
			//from grpc status
			if st, ok := status.FromError(err); ok {
				return int32(st.Code())
			}
			//from error coder implement
			if cd, ok := err.(coder); ok {
				return cd.Value()
			}
			err = errors.Unwrap(err)
		}
		return -1
	}
	return 0
}
