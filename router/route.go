package route

import (
	controller "go-challenge/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoute(e *echo.Echo) error {
	DeviceRoute(e)
	return nil
}

func DeviceRoute(e *echo.Echo) {
	g := e.Group("api")
	g.POST("/devices", controller.Store)
	g.GET("/devices/:id", controller.Show)
}
