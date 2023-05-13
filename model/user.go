package model

import (
	"net/http"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u User) Connect(r http.Request, rw http.ResponseWriter) {

}
