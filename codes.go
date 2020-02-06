package errors

import (
	"fmt"
)

type ErrCode int32

func (e ErrCode) Value() int32 {
	return int32(e)
}

func (e ErrCode) String() string {
	return fmt.Sprintf("ErrCode(%d)", e)
}

func (e ErrCode) Error() string {
	return e.String()
}

//ValueErr pure value error
func ValueErr(v int32) error {
	if v != 0 {
		return ErrCode(v)
	}
	return nil
}
