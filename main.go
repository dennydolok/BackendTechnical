package main

import (
	"dennydolok/experimentals-1/config"
	"dennydolok/experimentals-1/handlers/rest"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config := config.InitConfig()
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	rest.RegisterAPI(e, config)

	e.Start(":8080")

}
