package main

import (
    "fmt"
    "log"
    "net/url"

    "github.com/gorilla/websocket"
)

func main() {
    u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/echo"}

    conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatal("error de conexión:", err)
    }
    defer conn.Close()

    // Envía un mensaje al servidor
    message := "Hola, servidor"
    err = conn.WriteMessage(websocket.TextMessage, []byte(message))
    if err != nil {
        log.Println("error al enviar mensaje:", err)
        return
    }

    // Lee el mensaje de vuelta del servidor
    messageType, p, err := conn.ReadMessage()
    if err != nil {
        log.Println("error al leer mensaje:", err)
        return
    }
    fmt.Println("Mensaje enviado:", string(p), "Tipo:", messageType)
}

