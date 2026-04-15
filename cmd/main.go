package main

import (
	"log"

	"github.com/l10-bhushan/jwt_auth/internal/router"
)

func main() {
	log.Println("Server running on PORT :8080")
	cfg := router.Config{
		Addr: ":8080",
	}
	app := router.Application{
		Cfg: &cfg,
	}

	app.Run(app.Mount())

}
