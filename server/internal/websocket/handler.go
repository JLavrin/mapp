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
	Code  int    `json:"code"`
	Data  any    `json:"data"`
	Event string `json:"event"`
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

	b, err := json.Marshal(Res{
		Code: 200,
		Data: map[string]interface{}{
			"message": "Connected",
		},
		Event: "connected",
	})

	if err != nil {
		log.Println("[Websocket] MArachsal connected message")
	}

	err = conn.WriteMessage(websocket.TextMessage, b)

	if err != nil {
		log.Println("[Websocket] Error connected message")
	}

	for {
		select {
		case message := <-service.VehicleUpdates:
			b, err := json.Marshal(Res{
				Code:  200,
				Data:  message,
				Event: "vehicles_update",
			})
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
