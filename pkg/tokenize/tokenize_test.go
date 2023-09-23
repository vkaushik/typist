package tokenize

import (
	"reflect"
	"testing"
)

func TestSplitStringByWhitespaceOrPunctuation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "one",
			args: args{
				input: `Hello, world! This is a test string. It's separated by spaces and punctuation. 200-233-9843`,
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitStringByWhitespaceOrPunctuation(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitStringByWhitespaceOrPunctuation() = %v, want %v", got, tt.want)
			}
		})
	}
}
