package resp

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{name: "SimpleString", input: "+OK\r\n", expected: "OK"},
		{name: "Error", input: "-error message\r\n", expected: fmt.Errorf("error message")},
		{name: "Integer", input: ":100\r\n", expected: 100},
		{name: "BulkString", input: "$5\r\nhello\r\n", expected: []byte("hello")},
		{name: "NilBulkString", input: "$-1\r\n", expected: nil},
		{name: "Array", input: "*3\r\n+SET\r\n+mykey\r\n+myvalue\r\n", expected: []interface{}{"SET", "mykey", "myvalue"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := bufio.NewReader(bytes.NewReader([]byte(test.input)))
			decoded, err := Decode(reader)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if err, ok := test.expected.(error); ok {
				if decodedErr, ok := decoded.(error); !ok || decodedErr.Error() != err.Error() {
					t.Errorf("Decode(%s) = %v; want %v", test.input, decoded, test.expected)
				}
			} else {
				if !reflect.DeepEqual(decoded, test.expected) {
					t.Errorf("Decode(%s) = %v; want %v", test.input, decoded, test.expected)
				}
			}
		})
	}
}
