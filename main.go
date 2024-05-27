package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
)

type Params struct {
	LineId  string `json:"lineId"`
	RouteId string `json:"routeId"`
	Event   string `json:"event" required:"true"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func main() {
	flag.Parse()
	var addr = flag.String("addr", "localhost:8080", "http service address")
	var upgrader = websocket.Upgrader{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusInternalServerError)
			return
		}

		for {
			messageType, p, err := conn.ReadMessage()

			if err != nil {
				sendError(conn, err, 1)
				return
			}

			var params Params

			if err := json.Unmarshal(p, &params); err != nil {
				sendError(conn, err, 1)
				return
			}

			usedParams := ""

			val := reflect.ValueOf(params)
			for i := 0; i < val.NumField(); i++ {
				field := val.Field(i)
				fieldName := val.Type().Field(i).Name
				if field.String() != "" {
					usedParams += fieldName + "; "
				}
			}

			messageHandler(conn, p)
		}
	})

	fmt.Println("Server successfully started on", *addr)
	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		panic(err)
	}
}

func sendError(conn *websocket.Conn, err error, code int) {
	res := Error{
		Message: err.Error(),
		Code:    code,
	}

	msg, _ := json.Marshal(res)

	conn.WriteMessage(websocket.TextMessage, msg)
}

func messageHandler(conn *websocket.Conn, p []byte) {

}
