package vlc

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"gorchiver/lib/compression/vlc/table"
	"log"
	"strings"
	"unicode"
)

type EncoderDecoder struct {
	tableGenerator table.Generator
}

func NewEncoderDecoder(tableGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tableGenerator: tableGenerator}
}

// Encode encodes latin text (without symbols atm) into hex
func (ed EncoderDecoder) Encode(str string) []byte {
	table := ed.tableGenerator.NewTable(str)
	encoded := encodeBin(str, table)
	return buildEncodedFile(table, encoded)
}

func (ed EncoderDecoder) Decode(encodedData []byte) string {
	table, data := parseFile(encodedData)
	return table.Decode(data)
}

func parseFile(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBytesCount = 4
		dataSizeBytesCount  = 4
	)
	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	dataSizeBinary, data := data[:dataSizeBytesCount], data[dataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tableBinary, data := data[:tableSize], data[tableSize:]
	table := decodeTable(tableBinary)
	body := NewBinChunks(data).Join()

	return table, body[:dataSize]
}

func buildEncodedFile(table table.EncodingTable, data string) []byte {
	encodedTable := encodeTable(table)
	var buf bytes.Buffer
	buf.Write(encodeInt(len(encodedTable)))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTable)
	buf.Write([]byte(splitByChunks(data, chunkSize).Bytes()))
	return buf.Bytes()
}

func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))
	return res
}

func decodeTable(tableBinary []byte) table.EncodingTable {
	var table table.EncodingTable
	r := bytes.NewReader(tableBinary)
	if err := gob.NewDecoder(r).Decode(&table); err != nil {
		log.Fatal("can't serialize table: ", err)
	}
	return table
}

func encodeTable(table table.EncodingTable) []byte {
	var tableBuf bytes.Buffer
	if err := gob.NewEncoder(&tableBuf).Encode(table); err != nil {
		log.Fatal("can't serialize table: ", err)
	}
	return tableBuf.Bytes()
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
func encodeBin(str string, table table.EncodingTable) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(runeToBin(ch, table))
	}

	return buf.String()
}

// runeToBin function returns the string of binary sequence
// corresponding to the char, if it is in the table
func runeToBin(ch rune, table table.EncodingTable) string {
	res, ok := table[ch]
	if !ok {
		panic("unknown characker: " + string(ch))
	}
	return res
}
