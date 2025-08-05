package main

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	// upgrade connection to a websocket connection
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	go websocket.Writer(ws)
	// listen indefinitely for new messages through websocket connection
	websocket.Reader(ws)
}

func setupRoutes() {
	// map /ws to the serveWs function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
