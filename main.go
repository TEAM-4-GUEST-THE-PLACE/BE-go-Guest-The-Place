package main

import (
	"goGuestThePlace/config"
	"goGuestThePlace/migration"
	"goGuestThePlace/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	config.DBConfig()
	migration.RunMigration()
	routes.Routes(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start("localhost:5002"))
}
