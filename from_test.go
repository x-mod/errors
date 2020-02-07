package errors

import (
	"testing"
)

func TestCauseFrom(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"cause from origin",
			args{
				New("origin"),
			},
			false,
		},
		{
			"cause from error",
			args{
				Annotate(New("origin"), "cause"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unwrap(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValueFrom(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		{
			"normal error",
			fields{
				WithCode(New("error reason"), ErrNo(1)),
			},
			1,
		},
		{
			"code error",
			fields{
				ErrNo(2),
			},
			2,
		},
		{
			"no code error",
			fields{
				New("error"),
			},
			-1,
		},
		{
			"no error",
			fields{
				nil,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := tt.fields.err
			if got := ValueFrom(ce); got != tt.want {
				t.Errorf("ValueFrom error = %v, want %v", got, tt.want)
			}
		})
	}
}
