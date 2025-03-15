package main

import (
	"crud-with-cache/app/subscriber"
	"crud-with-cache/config"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config")
	}

	infra, err := subscriber.NewInfra(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize infra: %v", err))
	}
	server := subscriber.NewServer(infra)

	e := echo.New()
	server.RegisterRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
