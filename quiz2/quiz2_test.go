package quiz2

import "testing"

func TestFirstNonRepeatingChar(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{input: "Success"},
			want: "u",
		},
		{
			name: "Test case 2",
			args: args{input: "Help me!"},
			want: "H",
		},
		{
			name: "Test case 3",
			args: args{input: "The teacher"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstNonRepeatingChar(tt.args.input); got != tt.want {
				t.Errorf("FirstNonRepeatingChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
