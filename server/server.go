package server

import (
	"bufio"
	"fmt"
	"net"
)

func Run() {

	incoming_connections := make(chan net.Conn)

	go listen(incoming_connections)

	var connections []net.Conn

	incoming_messages := make(chan string)

	for {

		select {
		case conn := <-incoming_connections:
			{
				connections = append(connections, conn)

				go handle_connection(conn, incoming_messages)
			}
		case msg := <-incoming_messages:
			{
				fmt.Print(msg)

				for _, c := range connections {

					connection_writer := bufio.NewWriter(c)
					connection_writer.WriteString(msg)
					connection_writer.Flush()
				}
			}
		}

	}
}

func listen(incoming_connections chan<- net.Conn) {
	port := 30000

	listener, err := net.Listen("tcp", fmt.Sprint(":", port))

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("Listening on port", port)

	for {

		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		incoming_connections <- conn
	}
}

func handle_connection(c net.Conn, incoming_messages chan<- string) {

	fmt.Printf("connection new %s\n", c.RemoteAddr().String())

	defer func() {
		fmt.Printf("connection close %s\n", c.RemoteAddr().String())
		c.Close()
	}()

	writer := bufio.NewWriter(c)
	writer.WriteString("Hallo\n")

	reader := bufio.NewReader(c)

	for {
		message, err := reader.ReadBytes('\n')

		if err != nil {
			return
		}

		incoming_messages <- string(message)
	}
}
