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

func writeWs(ws *websocket.Conn, p []byte) (n int, err error) {
	err = ws.WriteMessage(websocket.TextMessage, p)

	if err != nil {
		return 0, err
	}

	return len(p), nil
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

	for ;; {
		_, msg, _ := conn.ReadMessage()

		log.Printf("adding message %s", string(msg))
		p.getStore().addMessage(string(msg))

		log.Printf("writing message %s", string(msg))
		_, err = writeWs(conn, msg)

		if err != nil {
			log.Panicf("webSocket write message %#v", err)
		}
	}
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
