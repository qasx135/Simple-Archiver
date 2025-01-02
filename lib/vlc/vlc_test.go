package vlc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	inputString := "Hello World"
	expected := "!hello !world"
	actual := prepareText(inputString)

	assert.Equal(t, expected, actual)
}