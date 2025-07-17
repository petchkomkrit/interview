package quiz1

import "testing"

func TestValidatePassword(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{input: "@Passw0rd"},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{input: "Passw0rd!"},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{input: "paSs#w0rd"},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{input: "Passw0rdx"},
			want: false,
		},
		{
			name: "Test case 5",
			args: args{input: "@Password"},
			want: false,
		},
		{
			name: "Test case 6",
			args: args{input: "@passw0rd"},
			want: false,
		},
		{
			name: "Test case 7",
			args: args{input: "@PASSW0RD"},
			want: false,
		},
		{
			name: "Test case 8",
			args: args{input: "@Pasw0rd"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidPassword(tt.args.input); got != tt.want {
				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
