package errors

import "testing"

func TestAnnotatef(t *testing.T) {
	type args struct {
		err    error
		format string
		args   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"annotationf",
			args{
				New("origin"),
				"format %v",
				[]interface{}{"hhh"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Annotatef(tt.args.err, tt.args.format, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Annotatef() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
