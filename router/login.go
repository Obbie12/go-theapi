package router

import (
	"go-theapi/controller"

	"github.com/labstack/echo/v4"
)

func InitLogin(e *echo.Echo) {
	e.POST("/login", controller.C_login_post)
}
