package models

import (
	"nhooyr.io/websocket"
)

type UserChannel chan *User

type User struct {
	Connection *websocket.Conn
	SeatNumber string
	Score      int
}

func NewUser(connection *websocket.Conn, seatNumber string) *User {
	return &User{
		Connection: connection,
		SeatNumber: seatNumber,
		Score:      0,
	}
}

func (u *User) AddClick() {
    u.Score++
}

type ByScore []*User

func (s ByScore) Len() int {
    return len(s)
}

func (s ByScore) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByScore) Less(i, j int) bool {
    return s[i].Score < s[j].Score
}
