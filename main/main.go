package main

import (
	"fmt"
	"os"

	"github.com/fogelfisch/hello/client"
	"github.com/fogelfisch/hello/server"
)

func main() {

	args := os.Args
	fmt.Println("args: ", args)

	if len(args) != 2 {
		fmt.Println("genau ein Argument ist anzugeben")
		panic(666)
	}

	if args[1] == "s" {
		server.Run()
	}

	if args[1] == "c" {
		client.Run()
	}
}
