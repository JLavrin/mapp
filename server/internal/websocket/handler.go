package websocket

import (
	"encoding/json"
	"github.com/JLavrin/mapp.git/server/internal/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Req struct {
	LineId string `json:"lineId"`
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
		select {
		case message := <-service.VehicleUpdates:
			b, err := json.Marshal(message)
			if err != nil {
				log.Println("[Websocket] Error marshalling message")
				continue
			}

			err = conn.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Println("[Websocket] Error sending message")
				return
			}
		}
	}
}
