package app

import (
	"crud-with-cache/config"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.Config{}

	infra := NewInfra(config)
	server := NewServer(infra)

	e := echo.New()
	server.RegisterRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
