package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunk string
type BinaryChunks []BinaryChunk
type HexChunk string
type HexChunks []HexChunk

const chunkSize = 8

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))
	for _, chunk := range bcs {
		res = append(res, chunk.ToHex())
	}
	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)

	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))
	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

func (hcs HexChunks) ToString() string {
	const sep = " "

	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder
	buf.WriteString(string(hcs[0]))
	for _, hc := range hcs[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(hc))
	}

	return buf.String()
}

// Encode encodes latin text (without symbols atm) into hex
func Encode(str string) string {
	str = prepareText(str)
	binStr := encodeBin(str)
	chunks := splitByChunks(binStr, chunkSize)

	return chunks.ToHex().ToString()
}

// prepareText prepares text to be fit for encode:
// changes upper case letters to: ! + lower case letter
// i.g.: My name is Ted -> !my name is !ted
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

// encodeBin encodes str into binary codes string without spaces
func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(runeToBin(ch))
	}

	return buf.String()
}

// runeToBin function returns the string of binary sequence
// corresponding to the char, if it is in the table
func runeToBin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown characker: " + string(ch))
	}

	return res
}

// getEncodingTable returns table of prepared values
// where each latin symbol has its own binary sequence
func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'c': "000101",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

// splitByChunks splits binary string by chunks with given size,
// i.g.: '100101011001010110010101' -> '10010101 10010101 10010101'
func splitByChunks(binStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(binStr)
	chunksCount := strLen / chunkSize

	if strLen / chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)
	var buf strings.Builder

	for i, ch := range binStr {
		buf.WriteString(string(ch))

		if (i+1) % chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	// if last chunk unfilled - fullfill it with 0's
	if buf.Len() != 0 {
		lastChank := buf.String()
		lastChank += strings.Repeat("0", chunkSize - len(lastChank))
		res = append(res, BinaryChunk(lastChank))
	}

	return res
}
