package http

import (
	"myProject/server/http/gateway"
	"sync"

	"github.com/labstack/echo/v4"
)

var server *echo.Echo
var once sync.Once

// New make new http server instance
func New() *echo.Echo {
	e := echo.New()

	return e
}

// Serve serve http server
func Serve() {
	once.Do(func() {
		server = New()
	})
	gateway.Register(server)
	err := server.Start(":8000")
	if err != nil {
		return
	}

}

// Server get http server instance
func Server() *echo.Echo {
	return server
}
