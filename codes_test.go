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
			if err := ErrNo(tt.args.v); Is(err, ErrNo(tt.args.v)) != tt.wantErr {
				t.Errorf("ErrNo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
