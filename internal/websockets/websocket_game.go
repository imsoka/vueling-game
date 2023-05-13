package websockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/imsoka/vueling-game/internal/models"
	"golang.org/x/text/cases"
)

var upGradeWebsocket = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: checkOrigin
}

func checkOrigin(r *http.Request) bool {
	//log.Printf()
	return r.Method == http.MethodGet
}

type Chanel struct {
	clickChannel ClickChanel
	userChanel UserChanel
}

type ClickChanel chan *models.Click
type UserChanel chan *models.User

type WebSocketGame struct {
	users map[string]*User
	joinGame UserChanel
	chanels *Chanel
}

func (w *WebSocketGame) UsersJoinManager() {
	for {
		select {
		case userJoin := <- w.joinGame:
		case click := <- w.chanels.clickChannel:
		}
	}
}