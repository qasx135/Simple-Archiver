package vlc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		// TODO: Add test cases.
		{
			name: "base test",
			bcs:  BinaryChunks{"00101111", "10000000"},
			want: "0010111110000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.bcs.Join(), "Join()")
		})
	}
}

func TestSplitByChunks(t *testing.T) {
	inputString := "00100000100110001001001"
	expected := BinaryChunks{"00100000", "10011000", "10010010"}
	actual := splitByChunks(inputString, 8)

	assert.Equal(t, expected, actual)
}

func TestNewBinChunks(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want BinaryChunks
	}{
		// TODO: Add test cases.
		{
			name: "base test",
			data: []byte{20, 30, 60, 18},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewBinChunks(tt.data), "NewBinChunks(%v)", tt.data)
		})
	}
}
