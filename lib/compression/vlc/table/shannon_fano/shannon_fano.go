package shannon_fano

import (
	"fmt"
	"gorchiver/lib/compression/vlc/table"
	"math"
	"sort"
	"strings"
)

type Generator struct{}

type charStat map[rune]int
type encodingTable map[rune]code

type code struct {
	Char     rune
	Quantity int
	Bits     uint32
	Size     int
}

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) NewTable(text string) table.EncodingTable {
	charStat := calcCharStat(text)
	return buildEncodingTable(charStat).Export()
}

func (et encodingTable) Export() map[rune]string {
	res := make(map[rune]string)

	for ch, code := range et {
		byteStr := fmt.Sprintf("%b", code.Bits)
		if lenDiff := code.Size - len(byteStr); lenDiff > 0 {
			byteStr = strings.Repeat("0", lenDiff) + byteStr
		}
		res[ch] = byteStr
	}

	return res
}

func buildEncodingTable(stat charStat) encodingTable {
	codes := make([]code, 0, len(stat))
	for ch, qty := range stat {
		codes = append(codes, code{
			Char:     ch,
			Quantity: qty,
		})
	}

	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity != codes[j].Quantity {
			return codes[i].Quantity > codes[j].Quantity
		}
		return codes[i].Char < codes[j].Char
	})

	assignCodes(codes)
	res := make(encodingTable)

	for _, code := range codes {
		res[code.Char] = code
	}

	return res
}

func assignCodes(codes []code) {
	if len(codes) < 2 {
		return
	}
	divider := bestDividerPosition(codes)
	for i := 0; i < len(codes); i++ {
		codes[i].Bits <<= 1
		codes[i].Size++

		if i >= divider {
			codes[i].Bits |= 1
		}
	}
	assignCodes(codes[:divider])
	assignCodes(codes[divider:])
}

// bestDividerPosition get sorted slice of codes
// and return position, than divides slice as balanced as possible
func bestDividerPosition(codes []code) int {
	total := 0
	for _, code := range codes {
		total += code.Quantity
	}

	left := 0
	prevDiff := math.MaxInt
	bestPosition := 0

	for i := 0; i < len(codes)-1; i++ {
		left += codes[i].Quantity
		right := total - left

		diff := abs(right - left)
		if diff >= prevDiff {
			break
		}
		prevDiff = diff
		bestPosition = i + 1
	}

	return bestPosition
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func calcCharStat(text string) charStat {
	res := make(charStat)
	for _, ch := range text {
		res[ch]++
	}
	return res
}
