package sockets

import (
	"fmt"
	"html/template"
	"net/http"
)

type PingServer struct {
	http.Handler
}

func (s *PingServer) pingHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "good")
}

func (p *PingServer) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("sockets.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

const jsonContentType = "application/json"

func NewPingServer() *PingServer {
	p := new(PingServer)

	router := http.NewServeMux()
	router.Handle("/ping", http.HandlerFunc(p.pingHandler))
	router.Handle("/index", http.HandlerFunc(p.indexHandler))

	p.Handler = router

	return p
}
