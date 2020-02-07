package errors

import (
	"errors"

	"testing"
)

func Test_annotator_Error(t *testing.T) {
	type fields struct {
		err        error
		annotation string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"message",
			fields{
				New("origin error"),
				"annotation",
			},
			"annotation origin error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Annotate(tt.fields.err, tt.fields.annotation)
			if got := err.Error(); got != tt.want {
				t.Errorf("annotator.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_annotator_Cause(t *testing.T) {
	type fields struct {
		err        error
		annotation string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"unwrap",
			fields{
				New("origin error"),
				"annotation comments",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Annotate(tt.fields.err, tt.fields.annotation)
			if err := errors.Unwrap(e); (err == tt.fields.err) != tt.wantErr {
				t.Errorf("errors.Unwrap error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAnnotate(t *testing.T) {
	type args struct {
		err        error
		annotation string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"normal",
			args{
				New("error"),
				"comment",
			},
			true,
		},
		{
			"abnormal",
			args{
				nil,
				"comment",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Annotate(tt.args.err, tt.args.annotation); (err != nil) != tt.wantErr {
				t.Errorf("Annotate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWithCode(t *testing.T) {
	type args struct {
		err  error
		code Code
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"normal",
			args{
				New("error"),
				ErrNo(1),
			},
			true,
		},
		{
			"abnormal",
			args{
				nil,
				ErrNo(2),
			},
			false,
		},
		{
			"abnormal",
			args{
				New("error"),
				nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WithCode(tt.args.err, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("WithCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCodeError(t *testing.T) {
	type args struct {
		code Code
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"normal",
			args{
				ErrNo(1),
			},
			true,
		},
		{
			"abnormal",
			args{
				nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CodeError(tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("CodeError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_codeErr_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"normal error",
			fields{
				WithCode(New("error reason"), ErrNo(2)),
			},
			"error reason",
		},
		{
			"code error",
			fields{
				ErrNo(2),
			},
			"ErrNo(2)",
		},
		{
			"code error",
			fields{
				ErrNo(3),
			},
			"ErrNo(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.err.Error(); got != tt.want {
				t.Errorf("codeErr.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithStack(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"stack error",
			args{
				New("origin"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WithStack(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("WithStack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"errorf",
			args{
				"format err: %v",
				[]interface{}{false},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Errorf(tt.args.format, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Errorf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
