package main

import (
	"fmt"
	"github.com/JLavrin/mapp.git/server/internal/util"
	"github.com/JLavrin/mapp.git/server/internal/websocket"
	"log"
	"net/http"
	"strconv"
)

func main() {
	util.LoadEnv()

	port := util.GetEnv("PORT", 8080)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("App running correctly."))
	})

	http.HandleFunc("/ws", websocket.Handler)

	portStr := strconv.Itoa(port)

	fmt.Printf("[Server started] available at http://localhost:%s\n", portStr)

	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}
