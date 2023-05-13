package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer conn.Close()

        // Ciclo de lectura de mensajes
        for {
            _, message, err := conn.ReadMessage()
            if err != nil {
                // Maneja el error de conexión cerrada
                if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
                    fmt.Println("Conexión cerrada por el cliente")
                } else {
                    fmt.Println("Error de lectura:", err)
                }
                break
            }
            fmt.Println("Mensaje recibido:", string(message))

            // Envía el mensaje de vuelta al cliente
            err = conn.WriteMessage(websocket.TextMessage, message)
            if err != nil {
                fmt.Println("Error de escritura:", err)
                break
            }
        }
    })

    fmt.Println("Servidor escuchando en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

