package main

import (
	"dennydolok/experimentals-1/config"
	"dennydolok/experimentals-1/handlers/rest"

	"github.com/labstack/echo/v4"
)

func main() {

	config := config.InitConfig()
	e := echo.New()

	rest.RegisterAPI(e, config)

	e.Start(":8080")

}
