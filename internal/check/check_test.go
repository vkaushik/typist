package check

import (
	"reflect"
	"testing"

	"github.com/vkaushik/typist/pkg/error"
)

func TestGetErrors(t *testing.T) {
	type args struct {
		master string
		test   string
	}
	tests := []struct {
		name  string
		args  args
		want  []error.TypingError
		want1 []error.TypingError
	}{
		{
			name: "1",
			args: args{
				master: "",
				test:   "",
			},
			want:  []error.TypingError{(error.TypingError{Remark: ""})},
			want1: []error.TypingError{(error.TypingError{Remark: ""})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetErrors(tt.args.master, tt.args.test)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetErrors() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetErrors() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// the word has two two
// the has has two

// hello world how
// hello how

// hello world. why
// hello why

// hello. .how
// be
