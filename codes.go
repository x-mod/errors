package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

//GRPCStatus make ErrNo support grpc status
func (e ErrNo) GRPCStatus() *status.Status {
	return status.New(codes.Code(e.Value()), e.Error())
}
