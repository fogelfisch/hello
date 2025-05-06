package server

import (
	"fmt"
	"net"
	"time"
)

func Server() {

	port := 30000

	listener, err := net.Listen("tcp", fmt.Sprint(":", port))

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("listening on", port, "...")

	conn, err := listener.Accept()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("new connection to", conn.RemoteAddr())

	conn.Write([]byte("Tschüss\n"))
	conn.Write([]byte("Du\n"))
	conn.Write([]byte("kleiner\n"))
	conn.Write([]byte("Sack\n"))

	time.Sleep(10 * time.Second)
}
