package sockets

import (
	"fmt"
	"net/http"
)

type PingServer struct {
	http.Handler
}

func (s PingServer) pingHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "good")
}

const jsonContentType = "application/json"

func NewPingServer() *PingServer {
	p := new(PingServer)

	router := http.NewServeMux()
	router.Handle("/ping", http.HandlerFunc(p.pingHandler))

	p.Handler = router

	return p
}
