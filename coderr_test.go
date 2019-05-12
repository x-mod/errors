package errors

import (
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
				code: &errorCode{1, "NumOneErr"},
				err:  New("origin"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := &codeErr{
				code: tt.fields.code,
				err:  tt.fields.err,
			}
			if err := ce.Cause(); (err != nil) != tt.wantErr {
				t.Errorf("codeErr.Cause() error = %v, wantErr %v", err, tt.wantErr)
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
				code: &errorCode{1, "NumOneErr"},
				err:  New("origin"),
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := &codeErr{
				code: tt.fields.code,
				err:  tt.fields.err,
			}
			if got := ce.Value(); got != tt.want {
				t.Errorf("codeErr.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
