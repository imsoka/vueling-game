package model

type User struct {
	ipAddress  string
	score      uint
	seat       string
	connection Connection
}

func NewUser(conn *Connection) *User {
	return &User{}
}
