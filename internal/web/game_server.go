package web

import (
	"log"
	"net/http"

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
	exist := gs.playerExists(seat)
	if !exist {
		p := models.NewUser(conn, seat)
		gs.Players = append(gs.Players, p)
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

func (gs *GameServer) ClickHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Printf("%v", err)
		return
	}
	for _, values := range r.PostForm {
		log.Printf("%v", values)
	}
}
