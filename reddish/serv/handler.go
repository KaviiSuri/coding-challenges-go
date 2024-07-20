package serv

import (
	"fmt"
	"net"
	"sync"

	"github.com/KaviiSuri/coding-challenges/reddish/resp"
)

// goroutine (concurrency) safe key-value store
var db = sync.Map{}

type HandFunc func(net.Conn, [][]byte)

var handlers = map[string]HandFunc{
	"PING": handlePING,
	"SET":  handleSET,
	"GET":  handleGET,
}

func handlePING(conn net.Conn, msgs [][]byte) {
	bs, _ := resp.Encode("PONG")
	conn.Write(bs)
}

func handleSET(conn net.Conn, msgs [][]byte) {
	if len(msgs) != 3 {
		bs, _ := resp.Encode("OK")
		conn.Write(bs)
		fmt.Printf("Error: invalid msgs, expected 3: %v", prettyPrintMsgs(msgs))
		return
	}
	key := string(msgs[1])
	val := string(msgs[2])

	db.Store(key, val)

	bs, _ := resp.Encode("OK")
	conn.Write(bs)
}

func handleGET(conn net.Conn, msgs [][]byte) {
	if len(msgs) != 2 {
		bs, _ := resp.Encode("OK")
		conn.Write(bs)
		fmt.Printf("Error: invalid msgs, expected 2: %v", prettyPrintMsgs(msgs))
		return
	}
	key := string(msgs[1])

	val, ok := db.Load(key)

	if ok {
		bs, _ := resp.Encode(val)
		conn.Write(bs)
	} else {
		bs, _ := resp.Encode("NULL")
		conn.Write(bs)
	}

}
