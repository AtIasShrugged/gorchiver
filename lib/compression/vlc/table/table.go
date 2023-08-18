package table

import "strings"

type Generator interface {
	NewTable(text string) EncodingTable
}

type EncodingTable map[rune]string

type decodingTree struct {
	Value string
	Left  *decodingTree
	Right *decodingTree
}

func (dt *decodingTree) add(code string, value rune) {
	currentNode := dt
	for _, ch := range code {
		switch ch {
		case '0':
			if currentNode.Left == nil {
				currentNode.Left = &decodingTree{}
			}
			currentNode = currentNode.Left
		case '1':
			if currentNode.Right == nil {
				currentNode.Right = &decodingTree{}
			}
			currentNode = currentNode.Right
		}
	}
	currentNode.Value = string(value)
}

func (dt *decodingTree) Decode(str string) string {
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

func (et EncodingTable) decodingTree() decodingTree {
	res := decodingTree{}
	for ch, code := range et {
		res.add(code, ch)
	}
	return res
}

func (et EncodingTable) Decode(str string) string {
	dt := et.decodingTree()

	return dt.Decode(str)
}
