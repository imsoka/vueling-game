package main

import (
    "net"
    "fmt"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        fmt.Println(err)
        return
    }

    message := string(buffer)
    fmt.Println("Message: ", message)
}
