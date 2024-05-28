package main

import (
	"fmt"
	"github.com/JLavrin/mapp.git/server/internal/websocket"
	"log"
	"net/http"
)

func main() {
	port :=

		http.HandleFunc("/ws", websocket.Handler)

	fmt.Printf("[Server started] available at http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
