package amh_go_utils

import (
	"errors"
	"testing"
)

type simpleStruct struct {
	msg     string
	wrapped error
}

func (err *simpleStruct) Error() string {
	return err.msg
}

func (err *simpleStruct) Unwrap() error {
	return err.wrapped
}

func TestDeadSimpleFormatter(t *testing.T) {
	type args struct {
		input error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple",
			args: args{errors.New("E1")},
			want: "* E1",
		},
		{
			name: "Simply wrapped",
			args: args{&simpleStruct{"E2", errors.New("E1")}},
			want: "* E2\n** E1",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeadSimpleFormatter(tt.args.input); got != tt.want {
				t.Errorf("[***]\nDeadSimpleFormatter() = \n%vWanted \n%v\n[***]", got, tt.want)
			}
		})
	}
}
