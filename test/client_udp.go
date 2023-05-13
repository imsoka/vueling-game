package main

import (
    "net"
    "fmt"
)

func main() {
    udpAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        fmt.Println(err)
        return
    }

    conn, err := net.DialUDP("udp", nil, udpAddr)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer conn.Close()

    message := "Hellow there udp server"
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Message sent", message)
}
