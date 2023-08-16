package vlc

import (
	"reflect"
	"testing"
)

func Test_encodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   encodingTable
		want DecodingTree
	}{
		{
			name: "Test 1",
			et: encodingTable{
				'a': "11",
				'b': "1001",
				's': "0101",
				'e': "101",
			},
			want: DecodingTree{
				Left: &DecodingTree{
					Right: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{
								Value: "s",
							},
						},
					},
				},
				Right: &DecodingTree{
					Left: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{
								Value: "b",
							},
						},
						Right: &DecodingTree{
							Value: "e",
						},
					},
					Right: &DecodingTree{
						Value: "a",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodingTable.DecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
