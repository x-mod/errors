package errors

import (
	"errors"

	"testing"
)

func Test_codeErr_Cause(t *testing.T) {
	type fields struct {
		code Code
		err  error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"cause",
			fields{
				code: ErrNo(1),
				err:  New("origin"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := WithCode(tt.fields.err, tt.fields.code)
			if err := errors.Unwrap(ce); (err != nil) != tt.wantErr {
				t.Errorf("errors.Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_codeErr_Value(t *testing.T) {
	type fields struct {
		code Code
		err  error
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		// TODO: Add test cases.
		{
			"value",
			fields{
				code: ErrNo(1),
				err:  New("origin"),
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := WithCode(tt.fields.err, tt.fields.code)
			if got := ValueFrom(ce); got != tt.want {
				t.Errorf("codeErr.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
