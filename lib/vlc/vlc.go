package vlc

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

type EncodingTable map[rune]string

const chunkSize = 8

func Encode(str string) string {
	// prepare text: M -> !m
	str = prepareText(str)
	// encode to binary: some text -> 01010101
	bStr := encodeBin(str)
	// split binary by chunks (8): bits to bytes -> '01010101 01010101 01010101'
	chunks := splitByChanks(bStr, chunkSize)
	fmt.Println(chunks)
	// bytes to hex -> '20 30 3C'ь
	// return hexChunksStr
	return ""
}

// prepareText prepares text to be fit for encode:
// changes upper case letters to: ! + lower case letter
// i.g.: My name is Michael -> !my name is !michael
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
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := encodingTable()
	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}

func encodingTable() EncodingTable {
	return EncodingTable {
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
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
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

// splitByChanks splits binary string by chunks with given size
// i.g. '100101011001010110010101' -> '10010101 10010101 10010101'
func splitByChanks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)

	chunksCount := strLen / chunkSize

	if strLen / chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder
	
	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i + 1) % chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize - len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}