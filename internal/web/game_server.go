package web

import (
    "log"
    "net/http"
    "encoding/json"

    "github.com/imsoka/vueling-game/internal/models"
    "nhooyr.io/websocket"
)

type GameServer struct {
    ServeMux http.ServeMux
    Players  []*models.User
}

func NewGameServer() *GameServer {
    gs := &GameServer{}

    gs.ServeMux.Handle("/", http.FileServer(http.Dir("./frontend")))
    gs.ServeMux.HandleFunc("/join", gs.JoinHandler)
    gs.ServeMux.HandleFunc("/click", gs.ClickHandler)

    return gs
}

func (gs *GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    gs.ServeMux.ServeHTTP(w, r)
}

func (gs *GameServer) JoinHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := websocket.Accept(w, r, nil)
    if err != nil {
        log.Printf("%v", err)
    }

    seat := r.URL.Query().Get("seat")
    
    if exist := gs.playerExists(seat); !exist {
        p := models.NewUser(conn, seat)
        gs.Players = append(gs.Players, p)
        log.Println("Player", seat, "joined the game")

    }
}

func (gs *GameServer) playerExists(seat string) bool {
    for _, v := range gs.Players {
        if v.SeatNumber == seat {
            return true
        }
    }
    return false
}

func (gs *GameServer) getPlayer(seat string) *models.User {
    for _, v := range gs.Players {
        if v.SeatNumber == seat {
            return v
        }
    }

    return nil
}

func (gs *GameServer) ClickHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
        return
    }

    defer r.Body.Close()

    var requestData struct {
        Msg string `json:"msg"`
        Seat string `json:"seat"`
    }

    log.Println(r.Body)
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        log.Printf("Error al decodificar el cuerpo de la solicitud: %v", err)
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }
    log.Printf("Datos recibidos: %+v", requestData)
    
    log.Println(requestData.Seat, "clicked")

    player := gs.getPlayer(requestData.Seat)
    if player == nil {
        // player := models.NewUser(conn, seat)
        // gs.Players = append(gs.Players, p)
        // log.Println("Player", seat, "joined the game")
        log.Println("Player", requestData.Seat, "does not exists, how is it playing?")
        return
    }

    player.AddClick()
    log.Println("Player", player.SeatNumber, "has", player.Score, "clicks")
}

