package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("Authorization")

		// auth to do
		if a == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(err, "test")
			return
		}

		defer conn.Close()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("received: %s", message)

			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println(err)
			}
		}
	})

	fmt.Printf("[] %s", "1")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
