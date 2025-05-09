package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Run(username string, server_ipadresse string) {

	port := 30000
	addr := fmt.Sprint(server_ipadresse, ":", port)

	fmt.Println("connecting to", addr, "...")

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		panic(err)
	}

	//	incoming_messages := make(chan string)
	fmt.Println("connected to", conn.RemoteAddr())

	go handle_connection(conn)

	console_reader := bufio.NewReader(os.Stdin)

	connection_writer := bufio.NewWriter(conn)

	fmt.Println("Simple Chat")
	fmt.Println("---------------------")

	for {
		// fmt.Print("-> ")
		message, err := console_reader.ReadString('\n')

		if err != nil {
			return
		}

		// convert CRLF to LF
		//	text = strings.Replace(text, "\n", "", -1)
		connection_writer.WriteString(username + ": " + message)
		connection_writer.Flush()
	}
}

func handle_connection(c net.Conn) {

	defer func() {
		fmt.Printf("connection close %s\n", c.RemoteAddr().String())
		c.Close()
	}()

	reader := bufio.NewReader(c)

	for {
		message, err := reader.ReadBytes('\n')

		if err != nil {
			return
		}

		fmt.Print(string(message))
	}
}
