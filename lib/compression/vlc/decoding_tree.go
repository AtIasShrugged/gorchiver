package vlc

import "strings"

type DecodingTree struct {
	Value string
	Left  *DecodingTree
	Right *DecodingTree
}

func (dt *DecodingTree) Add(code string, value rune) {
	currentNode := dt
	for _, ch := range code {
		switch ch {
		case '0':
			if currentNode.Left == nil {
				currentNode.Left = &DecodingTree{}
			}
			currentNode = currentNode.Left
		case '1':
			if currentNode.Right == nil {
				currentNode.Right = &DecodingTree{}
			}
			currentNode = currentNode.Right
		}
	}
	currentNode.Value = string(value)
}

func (dt *DecodingTree) Decode(str string) string {
	var buf strings.Builder

	currentNode := dt

	for _, ch := range str {
		if currentNode.Value != "" {
			buf.WriteString(currentNode.Value)
			currentNode = dt
		}
		switch ch {
		case '0':
			currentNode = currentNode.Left
		case '1':
			currentNode = currentNode.Right
		}
	}
	if currentNode.Value != "" {
		buf.WriteString(currentNode.Value)
	}
	return buf.String()
}

func (et encodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}
	for ch, code := range et {
		res.Add(code, ch)
	}
	return res
}
