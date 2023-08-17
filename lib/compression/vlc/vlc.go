package vlc

import (
	"strings"
	"unicode"
)
type EncoderDecoder struct {}

func NewEncoderDecoder() EncoderDecoder {
	return EncoderDecoder{}
}

type encodingTable map[rune]string

// Encode encodes latin text (without symbols atm) into hex
func (_ EncoderDecoder) Encode(str string) []byte {
	str = prepareText(str)
	chunks := splitByChunks(encodeBin(str), chunkSize)
	return chunks.Bytes()
}

func (_ EncoderDecoder) Decode(encodedData []byte) string {
	binString := NewBinChunks(encodedData).Join()
	decodingTree := getEncodingTable().DecodingTree()
	return recoverText(decodingTree.Decode(binString))
}

// prepareText prepares text to be fit for encode:
// changes uppercase letters to: ! + <lowercase letter>
// i.g.: My name is John -> !my name is !ted
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

// recoverText is opposite to prepareText: it return text
// to its original form with uppercase letter's
// by changing ! + <lowercase letter> to <uppercase letter>
// i.g.: !my name is !john -> My name is John
func recoverText(str string) string {
	var buf strings.Builder
	var isCapital bool
	for _, ch := range str {
		if isCapital && unicode.IsLetter(ch) {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false
			continue
		}
		if !isCapital && ch == '!' {
			isCapital = true
			continue
		}
		buf.WriteRune(ch)
	}
	if isCapital {
		buf.WriteRune('!')
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
