package main

import (
	"github.com/labstack/echo/v4"
	"github.com/x1bdev/gochat/internal/config"
	"github.com/x1bdev/gochat/internal/socket"
)

func main() {

	logger := config.NewLogger()
	logger.Setup()

	socket := socket.NewHandler()

	server := echo.New()
	server.GET("/ws", socket.HandleConnection)
	server.Start(":3000")
}
