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
		fmt.Println("Chat-Client starten mit: " + args[0] + " c <username> [<server_ipadresse>]")
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
			return
		}

		username := args[2]
		server_ipadresse := "127.0.0.1"

		if len(args) > 3 {
			server_ipadresse = args[3]
		}

		client.Run(username, server_ipadresse)
		return
	}

	hilfe_anzeigen()
}
