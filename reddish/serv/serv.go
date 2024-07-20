package serv

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/KaviiSuri/coding-challenges/reddish/resp"
)

func ProcessClient(conn net.Conn) {
	defer conn.Close()
	for {
		buffer := make([]byte, 1024)
		messageLen, err := conn.Read(buffer)
		if err != nil && messageLen == 0 {
			conn.Close()
			break
		}

		var bufferStr string = string(buffer[:messageLen])
		fmt.Printf("Received (%d): %s\n", messageLen, strings.ReplaceAll(bufferStr, "\r\n", "\\r\\n"))

		msgs, err := decodeRequest(bufferStr)
		if err != nil {
			fmt.Printf("Error Decoding Request: \n%v", err)
			conn.Close()
			break
		}

		fmt.Printf("Received: %v\n", string(msgs[0]))

		handler := handlers[string(msgs[0])]
		if handler == nil {
			fmt.Printf("No handler registered for %v", string(msgs[0]))
			bs, _ := resp.Encode("OK")
			conn.Write(bs)
		} else {
			handler(conn, msgs)
		}
	}
}

func decodeRequest(bufferStr string) ([][]byte, error) {
	decodedMsg, err := resp.Decode(bufio.NewReader(strings.NewReader(bufferStr)))
	if err != nil {
		return nil, fmt.Errorf("Error Parsing Message: %v", err)
	}
	items, ok := decodedMsg.([]any)
	if !ok {
		return nil, fmt.Errorf("Invalid Request: Non array value: %T \n%v", decodedMsg, resp.PrettyPrint(decodedMsg))
	}

	var msgs [][]byte
	for _, item := range items {
		byteSlice, ok := item.([]byte)
		if !ok {
			return nil, fmt.Errorf("Invalid Request: Non []byte value: %T \n%v", item, resp.PrettyPrint(item))
		}
		msgs = append(msgs, byteSlice)
	}
	return msgs, nil
}
