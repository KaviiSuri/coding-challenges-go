package bit

import (
	"bytes"
	"testing"
)

func TestBitWriter(t *testing.T) {
	tests := []struct {
		name     string
		bits     []bool
		expected []byte
	}{
		{
			name:     "Single bit - false",
			bits:     []bool{false},
			expected: []byte{0b00000000},
		},
		{
			name:     "Single bit - true",
			bits:     []bool{true},
			expected: []byte{0b10000000},
		},
		{
			name:     "Multiple bits - alternating",
			bits:     []bool{true, false, true, false, true, false, true, false},
			expected: []byte{0b10101010},
		},
		{
			name:     "Flush without complete byte",
			bits:     []bool{true, false, true},
			expected: []byte{0b10100000},
		},
		{
			name:     "Flush with multiple bytes",
			bits:     []bool{true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false},
			expected: []byte{0b10101010, 0b10101010},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			bw := NewBitWriter(&buf)

			for _, bit := range tt.bits {
				if err := bw.WriteBit(bit); err != nil {
					t.Fatalf("writeBit() error = %v", err)
				}
			}

			if err := bw.Flush(); err != nil {
				t.Fatalf("flush() error = %v", err)
			}

			if !bytes.Equal(buf.Bytes(), tt.expected) {
				t.Errorf("Expected %08b, but got %08b", tt.expected, buf.Bytes())
			}
		})
	}
}
