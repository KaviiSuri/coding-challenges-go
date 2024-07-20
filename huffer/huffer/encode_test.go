package huffer

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/KaviiSuri/coding-challenges/huffer/huffer/internal"
)

// Helper function to create a buffer with the encoded data
func encodeToBuffer(data []byte) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := Encode(buf, data)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected []byte // This should be the expected output after encoding
	}{
		{
			name:     "empty example",
			data:     []byte(""),
			expected: []byte{}, // Replace with actual expected encoded output
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf, err := encodeToBuffer(tt.data)
			if err != nil {
				t.Fatalf("Encode() error = %v", err)
			}
			if !bytes.Equal(buf.Bytes(), tt.expected) {
				t.Errorf("Encode() got = %v, want %v", buf.Bytes(), tt.expected)
			}
		})
	}
}

func TestWriteHeader(t *testing.T) {
	tests := []struct {
		name     string
		freqs    internal.Freqs
		wantErr  bool
		expected []byte
	}{
		{
			name:    "Normal case",
			freqs:   internal.Freqs{'a': 5, 'b': 3, 'c': 2},
			wantErr: false,
			expected: []byte{
				0x48, 0x46, 0x4D, 0x4E, // Magic number "HFMN"
				0x00, 0x00, 0x00, 0x03, // Number of characters (3)
				'a', 0x00, 0x00, 0x00, 0x05, // 'a': 5
				'b', 0x00, 0x00, 0x00, 0x03, // 'b': 3
				'c', 0x00, 0x00, 0x00, 0x02, // 'c': 2
			},
		},
		{
			name:    "Empty freqs",
			freqs:   internal.Freqs{},
			wantErr: false,
			expected: []byte{
				0x48, 0x46, 0x4D, 0x4E, // Magic number "HFMN"
				0x00, 0x00, 0x00, 0x00, // Number of characters (0)
			},
		},
		{
			name:    "Single character",
			freqs:   internal.Freqs{'x': 10},
			wantErr: false,
			expected: []byte{
				0x48, 0x46, 0x4D, 0x4E, // Magic number "HFMN"
				0x00, 0x00, 0x00, 0x01, // Number of characters (1)
				'x', 0x00, 0x00, 0x00, 0x0A, // 'x': 10
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := writeHeader(buf, tt.freqs)

			if (err != nil) != tt.wantErr {
				t.Errorf("WriteHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(buf.Bytes(), tt.expected) {
					t.Errorf("WriteHeader() = %v, want %v", buf.Bytes(), tt.expected)
				}

				// Additional check: read back and verify
				r := bytes.NewReader(buf.Bytes())
				err := ensureMagicNumber(r)
				if err != nil {
					t.Errorf("Failed to ensure magic number: %v", err)
				}
				readFreqs, err := decodeHeader(r)
				if err != nil {
					t.Errorf("Failed to read back header: %v", err)
				}
				if !reflect.DeepEqual(readFreqs, tt.freqs) {
					t.Errorf("Read back freqs = %v, want %v", readFreqs, tt.freqs)
				}
			}
		})
	}
}
