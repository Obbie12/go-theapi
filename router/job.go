package router

import (
	"go-theapi/controller"

	"github.com/labstack/echo/v4"
)

func Job(e *echo.Group) {
	e.GET("/jobs", controller.C_Get_Jobs)
	e.GET("/jobs/:id", controller.C_Get_JobById)
}
