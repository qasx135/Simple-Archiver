package vlc

import (
	"fmt"
	"strings"
	"unicode"
)

func Encode(str string) []byte {
	// prepare text: M -> !m
	str = prepareText(str)

	// encode to binary: some text -> 01010101
	bStr := encodeBin(str)

	// split binary by chunks (8): bits to bytes -> '01010101 01010101 01010101'
	chunks := splitByChunks(bStr, chunkSize)
	fmt.Println(chunks)

	// bytes to hex -> '20 30 3C'ь
	// return hexChunksStr
	return chunks.Bytes()
}

func Decode(encodedData []byte) string {
	// bChunks -> binary string
	bString := NewBinChunks(encodedData).Join()

	// build decoding tree
	dTree := encodingTable().DecodingTree()
	// bString (dTree) -> text
	// return decoded text
	return exportText(dTree.Decode(bString))
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
	return EncodingTable{
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

// i.g. !my name is !ted -> My name is Ted
func exportText(str string) string {
	var buf strings.Builder

	var isCapital bool

	for _, ch := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false
			continue
		}

		if ch == '!' {
			isCapital = true
			continue
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
