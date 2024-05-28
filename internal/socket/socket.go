package socket

import (
	"encoding/json"
	"io"
	"log/slog"
	"sync"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type SocketHandler struct {
	connections *sync.Map
	rooms       *sync.Map
}

func NewHandler() *SocketHandler {

	return &SocketHandler{
		connections: &sync.Map{},
		rooms:       &sync.Map{},
	}
}

func (s *SocketHandler) HandleConnection(c echo.Context) error {

	websocket.Handler(func(conn *websocket.Conn) {

		slog.Info("new incomming connection from", "addr", conn.RemoteAddr().String())
		defer conn.Close()

		s.connections.Store(conn, true)
		s.Listen(conn)

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func (s *SocketHandler) Listen(conn *websocket.Conn) {

	buffer := make([]byte, 1024)

	for {

		numberOfBytes, err := conn.Read(buffer)

		if err != nil {

			if err == io.EOF {
				break
			}

			slog.Error("could not read buffer", "error", err)
			continue
		}

		data := buffer[:numberOfBytes]
		message := &Message{}
		err = json.Unmarshal(data, message)

		if err != nil {
			slog.Error("could not parse message", "error", err)
			continue
		}

		s.Broadcast(conn, message)
	}
}

func (s *SocketHandler) Broadcast(conn *websocket.Conn, message *Message) {

	s.connections.Range(func(key, value any) bool {

		current := key.(*websocket.Conn)

		if conn == current {
			return true
		}

		data, err := json.Marshal(message)

		if err != nil {
			slog.Error("could not marshal message", "error", err)
			return true
		}

		slog.Info("message will be broadcast for listeners", "data", string(data))

		_, err = current.Write(data)

		if err != nil {
			slog.Error("could not send message", "error", err)
			return true
		}
		return true
	})
}
