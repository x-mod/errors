package errors

import "google.golang.org/grpc/status"

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
			cause, ok := err.(causer)
			if !ok {
				break
			}
			err = cause.Cause()
		}
		return -1
	}
	return 0
}

//CauseFrom get original error
func CauseFrom(err error) error {
	if err != nil {
		if cause, ok := err.(causer); ok {
			return cause.Cause()
		}
	}
	return err
}
