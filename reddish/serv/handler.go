package serv

import (
	"net"

	"github.com/KaviiSuri/coding-challenges/reddish/resp"
)

type HandFunc func(net.Conn, [][]byte)

var handlers = map[string]HandFunc{
	"PING": handlePING,
}

func handlePING(conn net.Conn, msgs [][]byte) {
	bs, _ := resp.Encode("PONG")
	conn.Write(bs)
}
