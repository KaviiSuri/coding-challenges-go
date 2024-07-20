package resp

import (
	"bytes"
	"fmt"
	"strconv"
)

// Encode takes a Go value and encodes it into RESP2 format.
func Encode(value interface{}) ([]byte, error) {
	var buf bytes.Buffer

	switch v := value.(type) {
	case string:
		buf.WriteByte(GetFirstByteFor(SimpleString))
		buf.WriteString(v)
		buf.WriteString("\r\n")
	case error:
		buf.WriteByte(GetFirstByteFor(SimpleError))
		buf.WriteString(v.Error())
		buf.WriteString("\r\n")
	case int:
		buf.WriteByte(GetFirstByteFor(Integer))
		buf.WriteString(strconv.Itoa(v))
		buf.WriteString("\r\n")
	case []byte:
		buf.WriteByte(GetFirstByteFor(BulkString))
		buf.WriteString(strconv.Itoa(len(v)))
		buf.WriteString("\r\n")
		buf.Write(v)
		buf.WriteString("\r\n")
	case nil:
		buf.WriteString("$-1\r\n")
	case []interface{}:
		buf.WriteByte(GetFirstByteFor(Array))
		buf.WriteString(strconv.Itoa(len(v)))
		buf.WriteString("\r\n")
		for _, elem := range v {
			elemBytes, err := Encode(elem)
			if err != nil {
				return nil, err
			}
			buf.Write(elemBytes)
		}
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}

	return buf.Bytes(), nil
}
