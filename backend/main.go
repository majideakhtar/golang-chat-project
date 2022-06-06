package main

import (
	"fmt"
	"net/http"

	"github.com/majideakhtar/golang-chat/pkg/websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket Endpoint reached")
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Majid's FullStack Chat Project using Go")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
