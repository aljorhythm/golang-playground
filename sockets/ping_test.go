package sockets

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	t.Run("GET /ping returns 200", func(t *testing.T) {
		server := NewPingServer()

		request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusOK)
	})

	t.Run("GET /index returns 200", func(t *testing.T) {
		server := NewPingServer()

		request, _ := http.NewRequest(http.MethodGet, "/index", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusOK)
	})

	t.Run("when we get a message over a websocket should be good", func(t *testing.T) {
		pingServer := NewPingServer()
		server := httptest.NewServer(pingServer)
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("could not open a ws connection on %s %v", wsURL, err)
		}
		defer ws.Close()

		message := "ping!"
		if err := ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			t.Fatalf("could not send message over ws connection %v", err)
		}

		time.Sleep(10 * time.Millisecond)

		store :=  pingServer.getStore()
		assert.Equal(t, []string{message},store.messages)
	})
}

