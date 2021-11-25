package sockets

import (
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
)

type PingStore struct {
	messages []string
}

func (store *PingStore) addMessage(message string) {
	store.messages = append(store.messages, message)
}

func NewStore() *PingStore {
	return &PingStore{[]string{}}
}

type PingServer struct {
	http.Handler
	store *PingStore
}

func (s *PingServer) getStore() *PingStore {
	return s.store
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

func (p *PingServer) webSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Panicf("webSocket %#v", err)
	}
	_, msg, _ := conn.ReadMessage()
	p.getStore().addMessage(string(msg))
}

const jsonContentType = "application/json"

func NewPingServer() *PingServer {
	p := new(PingServer)
	p.store = NewStore()
	router := http.NewServeMux()
	router.Handle("/ping", http.HandlerFunc(p.pingHandler))
	router.Handle("/index", http.HandlerFunc(p.indexHandler))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))

	p.Handler = router

	return p
}
