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
		want []byte
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := New()
			assert.Equalf(t, tt.want, encoder.Encode(tt.str), "Encode(%v)", tt.str)
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		encodedText []byte
		want        string
	}{
		{
			name:        "base test",
			encodedText: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want:        "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := New()
			assert.Equalf(t, tt.want, decoder.Decode(tt.encodedText), "Decode(%v)", tt.encodedText)
		})
	}
}
