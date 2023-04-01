package main

import (
	"fmt"
	"go-theapi/middlewares"
	"go-theapi/router"
	"go-theapi/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())

	config, err := util.LoadConfig(".")
	if err != nil {
		e.Logger.Fatal("cannot load config")
	}

	//create group
	a := e.Group("/auth")

	// set middleware
	middlewares.SetAuthMiddlewares(a)

	// Routes
	e.GET("/", hello)
	router.InitLogin(e)
	router.Job(a)

	// Start server
	port := fmt.Sprintf(":%s", config.ServerPort)
	e.Logger.Fatal(e.Start(port))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}
