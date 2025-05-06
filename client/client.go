package client

import (
	"bufio"
	"fmt"
	"net"
)

func Client() {

	port := 30000
	addr := fmt.Sprint(":", port)

	fmt.Println("connecting to", addr, "...")

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		panic(err)
	}

	fmt.Println("connected to", conn.RemoteAddr())

	close_connection := func() {
		fmt.Println("Closing connection...")
		conn.Close()
	}

	defer close_connection()

	reader := bufio.NewReader(conn)

	//buf := make([]byte, 0, 4096)

	//	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	for {
		// count, err := conn.Read(buf)

		buf, err := reader.ReadBytes('\n')

		if err != nil {
			return
		}

		fmt.Print(string(buf))
	}
}
