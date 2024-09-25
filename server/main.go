package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"time"
	"writedown-server/auth"
	"writedown-server/handler"
	"writedown-server/openapi"
	"writedown-server/repository"
)

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	server := handler.NewServer()

	err = repository.NewConnection(mysql.Config{
		User:      os.Getenv("DB_USERNAME"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOSTNAME") + ":" + os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
		Collation: "utf8mb4_general_ci",
		Loc:       jst,
	})
	if err != nil {
		log.Fatal(err)
	}

	store, err := repository.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	m := auth.NewMiddleware(swagger)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(store))
	e.Use(m.UserAuthMiddleware)

	openapi.RegisterHandlers(e, server)

	log.Fatal(e.Start(":8080"))
}
