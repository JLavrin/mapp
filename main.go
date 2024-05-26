package main

import (
	"flag"
	"net/http"
)

func main() {
	flag.Parse()
	var addr = flag.String("addr", "localhost:8080", "http service address")
	var upgrader = websocket.Upgrader{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Websocket endpoint"))
	})

	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		panic(err)
	}
}
