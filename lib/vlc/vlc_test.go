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

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Encode(tt.str), "Encode(%v)", tt.str)
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		encodedText string
		want        string
	}{
		{
			name:        "base test",
			encodedText: "20 30 3C 18 77 4A E4 4D 28",
			want:        "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Decode(tt.encodedText), "Decode(%v)", tt.encodedText)
		})
	}
}
