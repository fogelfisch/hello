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

	hilfe_anzeigen := func() {
		fmt.Println("Chat-Server starten mit: " + args[0] + " s")
		fmt.Println("Chat-Client starten mit: " + args[0] + " c <username>")
	}

	if len(args) < 2 {
		hilfe_anzeigen()
		return
	}

	if args[1] == "s" {
		server.Run()
		return
	}

	if args[1] == "c" {

		if len(args) < 3 {
			hilfe_anzeigen()
		}

		username := args[2]
		client.Run(username)
		return
	}

	hilfe_anzeigen()
}
