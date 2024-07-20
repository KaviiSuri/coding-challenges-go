package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"net/http"
	_ "net/http/pprof"

	"github.com/KaviiSuri/coding-challenges/reddish/serv"
)

const (
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = "6380"
	SERVER_TYPE = "tcp"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	fmt.Printf("Starting redis server on port %s ...\n", SERVER_PORT)
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening...", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Printf("Listening on %s ...\n", server.Addr())
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting...", err.Error())
			os.Exit(1)
		}
		go serv.ProcessClient(conn)
	}
}
