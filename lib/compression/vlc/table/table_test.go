package table

import (
	"reflect"
	"testing"
)

func TestEncodingTable_decodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   EncodingTable
		want decodingTree
	}{
		{
			name: "Test 1",
			et: EncodingTable{
				'a': "11",
				'b': "1001",
				's': "0101",
				'e': "101",
			},
			want: decodingTree{
				Left: &decodingTree{
					Right: &decodingTree{
						Left: &decodingTree{
							Right: &decodingTree{
								Value: "s",
							},
						},
					},
				},
				Right: &decodingTree{
					Left: &decodingTree{
						Left: &decodingTree{
							Right: &decodingTree{
								Value: "b",
							},
						},
						Right: &decodingTree{
							Value: "e",
						},
					},
					Right: &decodingTree{
						Value: "a",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.decodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodingTable.decodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
