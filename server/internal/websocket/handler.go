package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		m := "Error while upgrading connection"
		log.Printf("[Websocket] %s", m)
		http.Error(w, m, http.StatusBadRequest)
		return
	}

	handle(conn)

}

func handle(conn *websocket.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("[Websocket] Error while closing connection")
		}
	}()

}
