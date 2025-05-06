package main

import (
    "fmt"
    "net"
)




func main() {

    ln, err := net.Listen("tcp", ":8080")

    if err != nil {
        panic(err)
    }

    defer ln.Close()

    fmt.Println("Server is listening on port 8080...")

    conn, err
}
