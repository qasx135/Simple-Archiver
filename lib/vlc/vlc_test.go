package vlc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareText(t *testing.T) {
	inputString := "Hello World"
	expected := "!hello !world"
	actual := prepareText(inputString)

	assert.Equal(t, expected, actual)
}

func TestEncodeBin(t *testing.T) {
	inputString := "!lol"
	expected := "00100000100110001001001"
	actual := encodeBin(inputString)

	assert.Equal(t, expected, actual)
}

func TestSplitByChunks(t *testing.T) {
	inputString := "00100000100110001001001"
	expected := BinaryChunks{"00100000", "10011000", "10010010"}
	actual := splitByChunks(inputString, 8)

	assert.Equal(t, expected, actual)
}
