package spelling

import "testing"

func TestIsSpellingMistake(t *testing.T) {
	type args struct {
		expected string
		actual   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"example", "exaample"}, true},
		{"1", args{"example", "exaample"}, true},
		{"1", args{"example", "exemple"}, true},
		{"1", args{"spelling", "seepling"}, true},
		{"1", args{"spelling", "seplling"}, true},
		{"1", args{"spelling", "seepling"}, true},
		{"1", args{"spelling", "spelling"}, true},
		{"1", args{"spelling", "spelling"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSpellingMistake(tt.args.expected, tt.args.actual); got != tt.want {
				t.Errorf("IsSpellingMistake() = %v, want %v, args: %v", got, tt.want, tt.args)
			}
		})
	}
}
