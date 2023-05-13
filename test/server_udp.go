package main

import (
    "net"
    "fmt"
)

func main() {
    udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        fmt.Println(err)
        return
    }

    conn, err := net.ListenUDP("udp", udpAddr)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer conn.Close()

    for {
        buffer := make([]byte, 1024)
        _, addr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println(err)
            continue
        }

        message := string(buffer)
        fmt.Println("Mensaje de", addr.String(), ":", message)
    }
}
