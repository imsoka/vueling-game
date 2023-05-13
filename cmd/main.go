package main

import (
	"log"
	"net/http"

	"github.com/imsoka/vueling-game/api"
)

func main() {
	api.SetRoutes()
	log.Fatal(http.ListenAndServe(":80", nil))
}

func setupAPI() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}
