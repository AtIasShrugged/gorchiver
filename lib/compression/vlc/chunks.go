package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunk string
type BinaryChunks []BinaryChunk

const chunkSize = 8

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))
	for _, code := range data {
		res = append(res, NewBinChunk(code))
	}
	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

// Join joins chunks into one line and returns as string
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder
	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}
	return buf.String()
}

func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))
	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}
	return res
}

func (bc BinaryChunk) Byte() byte {
	res, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}
	return byte(res)
}

// splitByChunks splits binary string by chunks with given size,
// i.g.: '100101011001010110010101' -> '10010101 10010101 10010101'
func splitByChunks(binStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(binStr)
	chunksCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)
	var buf strings.Builder

	for i, ch := range binStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	// if last chunk unfilled - fullfill it with 0's
	if buf.Len() != 0 {
		lastChank := buf.String()
		lastChank += strings.Repeat("0", chunkSize-len(lastChank))
		res = append(res, BinaryChunk(lastChank))
	}

	return res
}
