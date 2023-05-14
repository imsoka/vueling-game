package web

import (
    "log"
    "net/http"
    "encoding/json"
    "sort"

    "github.com/imsoka/vueling-game/internal/models"
    "nhooyr.io/websocket"
)

type GameServer struct {
    ServeMux    http.ServeMux
    Players     []*models.User
}

func NewGameServer() *GameServer {
    gs := &GameServer{}

    gs.ServeMux.Handle("/", http.FileServer(http.Dir("./frontend")))
    gs.ServeMux.HandleFunc("/join", gs.JoinHandler)
    gs.ServeMux.HandleFunc("/click", gs.ClickHandler)
    gs.ServeMux.HandleFunc("/update", gs.UpdateHandler)

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

    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        log.Printf("Error al decodificar el cuerpo de la solicitud: %v", err)
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }
    
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

func (gs *GameServer) UpdateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
        return
    }

    defer r.Body.Close()

    var requestData struct {
        Msg     string  `json:"msg"`
        Seat    string  `json:"seat"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        log.Printf("Error al decodificar el cuerpo de la solicitud: %v", err)
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }

    if exists := gs.playerExists(requestData.Seat); !exists {
        log.Println("Player", requestData.Seat, "does not exists, how is it playing?")
        return
    }

    sort.Sort(models.ByScore(gs.Players))

    type RelPlayer struct {
        Seat    string  `json:"seat"`
        Score   int     `json:"score"`
    }

    var responseData struct {
        Prev RelPlayer `json:prev"`
        Next RelPlayer `json:next"`
    }

    var currentPlayer *models.User
    var prevPlayer *models.User
    var nextPlayer *models.User

    for idx, p := range gs.Players {
        if p.SeatNumber == requestData.Seat {
            currentPlayer = p
            if idx > 0 {
                prevPlayer = gs.Players[idx - 1]
            }
            if idx < len(gs.Players) - 1 {
                nextPlayer = gs.Players[idx + 1]
            }
            break
        }
    }

    if currentPlayer == nil {
        log.Println("Error while searching current player")
        return
    }

    if prevPlayer != nil {
        responseData.Prev.Seat = prevPlayer.SeatNumber
        responseData.Prev.Score = prevPlayer.Score
    } else {
        responseData.Prev.Seat = "-1"
        responseData.Prev.Score = -1
    }

    if nextPlayer != nil {
        responseData.Next.Seat = nextPlayer.SeatNumber
        responseData.Next.Score = nextPlayer.Score
    } else {
        responseData.Next.Seat = "-1"
        responseData.Next.Score = -1
    }

    log.Println(responseData)

    responseBytes, err := json.Marshal(responseData)
    if err != nil {
        log.Println("Error while encoding response", err)
        return
    }

    // if err := currentPlayer.Connection.Write(r.Context(), websocket.MessageText, responseBytes); err != nil {
    //     log.Println("Error sending response to client", err)
    //     return
    // }
    w.Header().Set("Content-Type", "application/json")
    w.Write(responseBytes)
}
