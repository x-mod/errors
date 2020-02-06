package errors

import "testing"

func TestValueErr(t *testing.T) {
	type args struct {
		v int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"value error",
			args{
				1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ErrCode(tt.args.v); Is(err, ErrCode(tt.args.v)) != tt.wantErr {
				t.Errorf("ValueErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
