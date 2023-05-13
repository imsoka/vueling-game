package main

import (
    "net"
    "fmt"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    message := "Hello there"
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Message sent: ", message)
}
