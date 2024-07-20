package serv

import "strings"

func prettyPrintMsgs(msgs [][]byte) string {
	ss := []string{}
	for _, msg := range msgs {
		ss = append(ss, string(msg))
	}
	return strings.Join(ss, " ")
}
