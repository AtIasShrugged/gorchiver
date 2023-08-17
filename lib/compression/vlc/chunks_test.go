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
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: "010111110000000",
		},
		{
			name: "Test 2",
			bcs:  BinaryChunks{""},
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

func TestBinaryChunk_Byte(t *testing.T) {
	tests := []struct {
		name string
		bc   BinaryChunk
		want byte
	}{
		{
			name: "Test 1",
			bc:   "00101010",
			want: 42,
		},
		{
			name: "Test 2",
			bc:   "00000000",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.Byte(); got != tt.want {
				t.Errorf("BinaryChunk.Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_Bytes(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want []byte
	}{
		{
			name: "Test 1",
			bcs:  BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
			want: []byte{20, 30, 60, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryChunks.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinChunks(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want BinaryChunks
	}{
		{
			name: "Test 1",
			data: []byte{20, 30, 60, 18},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinChunks(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
