package huffer

import (
	"bytes"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		// {
		// 	name: "Basic test with small data",
		// 	data: []byte("hello world"),
		// },
		// {
		// 	name: "Test with empty data",
		// 	data: []byte{},
		// },
		{
			name: "Test with repeated data",
			data: bytes.Repeat([]byte("a"), 1000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			// Encode data
			if err := Encode(&buf, tt.data); err != nil {
				t.Fatalf("Encode() error = %v", err)
			}

			// Decode data
			got, err := Decode(&buf)
			if err != nil {
				t.Fatalf("Decode() error = %v", err)
			}

			// Check if decoded data matches original data
			if !bytes.Equal(got, tt.data) {
				t.Errorf("Decode() = %v, want %v", string(got), string(tt.data))
			}
		})
	}
}
