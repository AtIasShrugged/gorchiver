package shannon_fano

import (
	"reflect"
	"testing"
)

func Test_bestDividerPosition(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "one element",
			codes: []code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "two elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "uncertainty (need rightmost) three elements",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "uncertainty (need rightmost) four elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "many elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDividerPosition(tt.codes); got != tt.want {
				t.Errorf("bestDividerPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name: "two elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "three elements, certain position",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1}, // 0 in binary
				{Quantity: 1, Bits: 2, Size: 2}, // 10 in binary
				{Quantity: 1, Bits: 3, Size: 2}, // 11 in binary
			},
		},
		{
			name: "three elements, uncertain position",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 1, Bits: 0, Size: 1}, // 0 in binary
				{Quantity: 1, Bits: 2, Size: 2}, // 10 in binary
				{Quantity: 1, Bits: 3, Size: 2}, // 11 in binary
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)

			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("got : %v, want: %v", tt.codes, tt.want)
			}
		})
	}
}

func Test_buildEncodingTable(t *testing.T) {
	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "Test 1",
			text: "abbbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Size:     2,
					Bits:     3,
				},
				'b': code{
					Char:     'b',
					Quantity: 3,
					Size:     1,
					Bits:     0,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Size:     2,
					Bits:     2,
				},
			},
		},
		{
			name: "Test 2",
			text: "aaccbb",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 2,
					Size:     1,
					Bits:     0,
				},
				'b': code{
					Char:     'b',
					Quantity: 2,
					Size:     2,
					Bits:     2,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Size:     2,
					Bits:     3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildEncodingTable(calcCharStat(tt.text)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildEncodingTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
