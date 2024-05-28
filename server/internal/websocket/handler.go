package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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

	for {
		_, message, err := conn.ReadMessage()
		start := time.Now()

		if err != nil {
			log.Println("[Websocket] Message read error")

			r := Res{
				Code:    400,
				Message: "Message decode error",
			}

			b := toJson(&r)

			conn.WriteMessage(websocket.TextMessage, b)
			return
		}

		end := time.Now()
		r := []byte(fmt.Sprintf("U sent: %s => %s", message, end.Sub(start)))
		conn.WriteMessage(websocket.TextMessage, r)
	}
}

func toJson(msg *Res) []byte {
	b, err := json.Marshal(msg)

	if err != nil {
		log.Printf("[toJson] Enconding error")
		return nil
	}

	return b
}
