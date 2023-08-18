package vlc

import (
	"testing"
)

func Test_recoverText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Test 1",
			str:  "!my name is !john",
			want: "My name is John",
		},
		{
			name: "Test 2",
			str:  "!a!a!a!a",
			want: "AAAA",
		},
		{
			name: "Test 3",
			str:  "",
			want: "",
		},
		{
			name: "Test 4",
			str:  "   ",
			want: "   ",
		},
		{
			name: "Test 5",
			str:  "!!no!t !!r!",
			want: "!NoT !R!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recoverText(tt.str); got != tt.want {
				t.Errorf("recoverText() = %v, want %v", got, tt.want)
			}
		})
	}
}
