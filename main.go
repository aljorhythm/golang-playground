package main

import (
	"github.com/aljorhythm/golang-playground/sockets"
	"log"
	"net/http"
)

func main() {
	pingServer := sockets.NewPingServer()
	httpServer := http.NewServeMux()
	httpServer.Handle("/", pingServer.Handler)
	log.Println("Running server")
	if err := http.ListenAndServe(":80", httpServer); err != nil {
		log.Fatalf("%#v", err)
	}
}
