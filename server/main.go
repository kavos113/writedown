package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"writedown-server/openapi"
)

func main() {
	server := openapi.NewServer()

	e := echo.New()

	openapi.RegisterHandlers(e, server)

	log.Fatal(e.Start("localhost:8080"))
}
