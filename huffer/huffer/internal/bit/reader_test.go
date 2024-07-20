package bit

import (
	"bytes"
	"reflect"
	"testing"
)

func TestBitReader(t *testing.T) {
	tests := []struct {
		data []byte
		bits []bool
	}{
		{
			data: []byte{0b10101010},
			bits: []bool{true, false, true, false, true, false, true, false},
		},
		{
			data: []byte{0b11111111},
			bits: []bool{true, true, true, true, true, true, true, true},
		},
		{
			data: []byte{0b00000000},
			bits: []bool{false, false, false, false, false, false, false, false},
		},
		{
			data: []byte{0b10101010, 0b00000000},
			bits: []bool{true, false, true, false, true, false, true, false, false, false, false, false, false, false, false, false},
		},
	}

	for _, tt := range tests {
		reader := NewBitReader(bytes.NewReader(tt.data))
		output, err := reader.ReadAllBits()
		if err != nil {
			t.Fatalf("ReadBit() error = %v", err)
		}
		if !reflect.DeepEqual(output, tt.bits) {
			t.Errorf("for %08b\n\nReadBit() = %v,\nWantBit() = %v", tt.data, output, tt.bits)
		}
	}
}
