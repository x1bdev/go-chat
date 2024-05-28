package socket

import (
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
			slog.Error("could not read buffer, error: %v", err)
			return
		}

		_ = numberOfBytes

	}
}
