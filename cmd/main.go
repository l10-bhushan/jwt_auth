package main

import (
	"log"

	"github.com/l10-bhushan/jwt_auth/internal/router"
)

func main() {
	log.Println("Jwt auth example")
	cfg := router.Config{
		Addr: ":8080",
	}
	app := router.Application{
		Cfg: &cfg,
	}
	app.Run(app.Mount())
}
