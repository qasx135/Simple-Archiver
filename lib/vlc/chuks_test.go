package vlc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitByChunks(t *testing.T) {
	inputString := "00100000100110001001001"
	expected := BinaryChunks{"00100000", "10011000", "10010010"}
	actual := splitByChunks(inputString, 8)

	assert.Equal(t, expected, actual)
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.bcs.ToHex(), "ToHex()")
		})
	}
}
