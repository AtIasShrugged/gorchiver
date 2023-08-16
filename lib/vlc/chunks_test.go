package vlc

import (
	"reflect"
	"testing"
)

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

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		{
			name: "Test 1",
			bcs: BinaryChunks{"0101111", "10000000"},
			want: "010111110000000",
		},
		{
			name: "Test 2",
			bcs: BinaryChunks{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("BinaryChunks.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunk_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bc   BinaryChunk
		want HexChunk
	}{
		{
			name: "Test 1",
			bc:   "11110111",
			want: "F7",
		},
		{
			name: "Test 2",
			bc:   "11111111",
			want: "FF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.ToHex(); got != tt.want {
				t.Errorf("BinaryChunk.ToHex() = %v, want %v", got, tt.want)
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

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "Test 1",
			str:  "20 30 3C 18 77 4A E4 4D 28",
			want: HexChunks{"20", "30", "3C", "18", "77", "4A", "E4", "4D", "28"},
		},
		{
			name: "Test 2",
			str:  "",
			want: HexChunks{""},
		},
		{
			name: "Test 3",
			str:  "10",
			want: HexChunks{"10"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hc   HexChunk
		want BinaryChunk
	}{
		{
			name: "Test 1",
			hc:   HexChunk("2A"),
			want: BinaryChunk("00101010"),
		},
		{
			name: "Test 2",
			hc:   HexChunk("2F"),
			want: BinaryChunk("00101111"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.ToBinary(); got != tt.want {
				t.Errorf("HexChunk.ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hcs  HexChunks
		want BinaryChunks
	}{
		{
			name: "Test 1",
			hcs:  HexChunks{"0F", "88"},
			want: BinaryChunks{"00001111", "10001000"},
		},
		{
			name: "Test 2",
			hcs:  HexChunks{"FA", "AA", "00", "42"},
			want: BinaryChunks{"11111010", "10101010", "00000000", "01000010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexChunks.ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
