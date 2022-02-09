package gateway

import "github.com/labstack/echo/v4"

// Register register all modules to router
func Register(server *echo.Echo) {
	RegisterApplication(server)
}
