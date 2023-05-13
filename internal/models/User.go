package models

import (
	"nhooyr.io/websocket"
)

type UserChannel chan *User

type User struct {
	Connection *websocket.Conn
	SeatNumber string
	Score      uint
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
