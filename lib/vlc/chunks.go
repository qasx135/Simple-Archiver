package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

type EncodingTable map[rune]string

type HexChunk string

type HexChunks []HexChunk

const chunkSize = 8

const hexChunkSeparator = " "

// splitByChunks splits binary string by chunks with given size
// i.g. '100101011001010110010101' -> '10010101 10010101 10010101'
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)

	chunksCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		// chunk -> hexChunk
		hexChunk := chunk.ToHex()
		res = append(res, hexChunk)
	}

	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic("cant parse binary chunk: " + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

func (hcs HexChunks) ToString() string {
	// 20 30 3C
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}
	var buf strings.Builder
	buf.WriteString(string(hcs[0]))
	for _, chunk := range hcs[1:] {
		buf.WriteString(hexChunkSeparator)
		buf.WriteString(string(chunk))
	}
	return buf.String()
}

func NewHexChunks(str string) HexChunks {
	parts := strings.Split(str, hexChunkSeparator)
	res := make(HexChunks, 0, len(parts))
	for _, part := range parts {
		res = append(res, HexChunk(part))
	}

	return res
}

func (hcs HexChunks) ToBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(hcs))

	for _, chunk := range hcs {
		bChunk := chunk.ToBinary()
		res = append(res, bChunk)
	}
	return res
}

func (hc HexChunk) ToBinary() BinaryChunk {
	num, err := strconv.ParseUint(string(hc), 16, chunkSize)
	if err != nil {
		panic("cant parse hex chunk: " + err.Error())
	}
	res := fmt.Sprintf("%08b", num)
	return BinaryChunk(res)
}

func (bcs BinaryChunks) Join() string {
	var buf strings.Builder
	for _, chunk := range bcs {
		buf.WriteString(string(chunk))
	}
	return buf.String()
}
