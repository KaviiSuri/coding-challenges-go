package resp

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{name: "SimpleString", input: "OK", expected: "+OK\r\n"},
		{name: "Error", input: fmt.Errorf("error message"), expected: "-error message\r\n"},
		{name: "Integer", input: 100, expected: ":100\r\n"},
		{name: "BulkString", input: []byte("hello"), expected: "$5\r\nhello\r\n"},
		{name: "NilBulkString", input: nil, expected: "$-1\r\n"},
		{name: "Array", input: []interface{}{"SET", "mykey", "myvalue"}, expected: "*3\r\n+SET\r\n+mykey\r\n+myvalue\r\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encoded, err := Encode(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(encoded) != test.expected {
				t.Errorf("Encode(%v) = %s; want %s", test.input, encoded, test.expected)
			}
		})
	}
}
