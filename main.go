package main
import (
	"fmt"
	"goGuestThePlace/config"
)

func main() {
	e := echo.New()
	// Tambahin migration disini
	
	config.DBConfig()
	// Tambahin Route disini

	e.Logger.Fatal(e.Start(":8000"))
}