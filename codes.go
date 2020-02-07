package errors

import (
	"fmt"
)

type ErrNo int32

func (e ErrNo) Value() int32 {
	return int32(e)
}

func (e ErrNo) String() string {
	return fmt.Sprintf("ErrNo(%d)", e)
}

func (e ErrNo) Error() string {
	return e.String()
}
