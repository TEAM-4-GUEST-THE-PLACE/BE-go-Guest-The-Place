package main

import (
	"goGuestThePlace/config"
	"goGuestThePlace/migration"
	"goGuestThePlace/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.DBConfig()
	migration.RunMigration()
	routes.Routes(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start("localhost:5002"))

}
