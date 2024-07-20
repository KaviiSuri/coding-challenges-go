package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Decode takes a RESP2 encoded byte slice and decodes it into a Go value.
func Decode(reader *bufio.Reader) (interface{}, error) {
	prefix, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	t, ok := GetDataType(prefix)
	if !ok {
		return nil, fmt.Errorf("unknown prefix: %c", prefix)
	}

	switch t {
	case SimpleString:
		line, err := readLine(reader)
		if err != nil {
			return nil, err
		}
		return line, nil
	case SimpleError:
		line, err := readLine(reader)
		if err != nil {
			return nil, err
		}
		return fmt.Errorf(line), nil
	case Integer:
		line, err := readLine(reader)
		if err != nil {
			return nil, err
		}
		return strconv.Atoi(line)
	case BulkString:
		line, err := readLine(reader)
		if err != nil {
			return nil, err
		}
		length, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		if length == -1 {
			return nil, nil
		}
		buf := make([]byte, length)
		if _, err := io.ReadFull(reader, buf); err != nil {
			return nil, err
		}
		if _, err := reader.Discard(2); err != nil {
			return nil, err
		}
		return buf, nil
	case Array:
		line, err := readLine(reader)
		if err != nil {
			return nil, err
		}
		length, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		array := make([]interface{}, length)
		for i := 0; i < length; i++ {
			elem, err := Decode(reader)
			if err != nil {
				return nil, err
			}
			array[i] = elem
		}
		return array, nil
	default:
		return nil, fmt.Errorf("unknown prefix: %c", prefix)
	}
}

// PrettyPrint takes a decoded RESP2 value and returns a formatted string.
func PrettyPrint(value interface{}) string {
	var sb strings.Builder
	prettyPrintValue(value, &sb, 0)
	return sb.String()
}

func prettyPrintValue(value interface{}, sb *strings.Builder, indent int) {
	indentStr := strings.Repeat("  ", indent)

	switch v := value.(type) {
	case string:
		sb.WriteString(indentStr + fmt.Sprintf("String: \"%s\"\n", v))
	case error:
		sb.WriteString(indentStr + fmt.Sprintf("Error: %s\n", v.Error()))
	case int:
		sb.WriteString(indentStr + fmt.Sprintf("Integer: %d\n", v))
	case []byte:
		sb.WriteString(indentStr + fmt.Sprintf("BulkString: \"%s\"\n", string(v)))
	case nil:
		sb.WriteString(indentStr + "Nil\n")
	case []interface{}:
		sb.WriteString(indentStr + "Array:\n")
		for _, item := range v {
			prettyPrintValue(item, sb, indent+1)
		}
	default:
		sb.WriteString(indentStr + fmt.Sprintf("Unknown Type: %v\n", v))
	}
}

// readLine reads a line (terminated by \r\n) from the reader.
func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	// Remove \r\n
	return string(line[:len(line)-2]), nil
}
