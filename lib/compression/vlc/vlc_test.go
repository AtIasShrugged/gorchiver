package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Test 1",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "Test 2",
			str:  "AAAA",
			want: "!a!a!a!a",
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
			str:  "!NoT !R!",
			want: "!!no!t !!r!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Test 1",
			str:  " ",
			want: "11",
		},
		{
			name: "Test 2",
			str:  "  ",
			want: "1111",
		},
		{
			name: "Test 3",
			str:  "",
			want: "",
		},
		{
			name: "Test 4",
			str:  "qwe",
			want: "0000000000010000011101",
		},
		{
			name: "Test 5",
			str:  "!hello !world",
			want: "001000001110100100100100110001110010000000011100010100000100100101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "Test 1",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := NewEncoderDecoder()
			if got := encoder.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		encodedData []byte
		want        string
	}{
		{
			name:        "Test 1",
			encodedData: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want:        "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := NewEncoderDecoder()
			if got := decoder.Decode(tt.encodedData); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
