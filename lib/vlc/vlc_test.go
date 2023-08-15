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

func Test_splitByChunks(t *testing.T) {
	type args struct {
		binStr    string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "Test 1",
			args: args{binStr: "100101011001010110010101", chunkSize: 8},
			want: BinaryChunks{"10010101", "10010101", "10010101"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.binStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "Test 1",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryChunks.ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str string
		want string
	}{
		{
			name: "main test",
			str: "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
