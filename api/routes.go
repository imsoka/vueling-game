package api

import "net/http"

func SetRoutes() {
	//Main route, front page, opens websocket for client
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	//Get users scores

	//Click the button, update score
}
