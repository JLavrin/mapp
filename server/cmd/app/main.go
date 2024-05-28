package main

import (
	"fmt"
	"github.com/JLavrin/mapp.git/server/internal/util"
	"github.com/JLavrin/mapp.git/server/internal/websocket"
	"log"
	"net/http"
)

func main() {
	port := util.GetEnv("PORT", ":8081")

	http.HandleFunc("/ws", websocket.Handler)

	fmt.Printf("[Server started] available at http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
