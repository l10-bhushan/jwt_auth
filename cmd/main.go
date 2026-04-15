package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/l10-bhushan/jwt_auth/internal/router"
)

func main() {
	log.Println("Server running on PORT :8080")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables...")
	}

	addr := os.Getenv("PORT")

	cfg := router.Config{
		Addr: addr,
	}
	app := router.Application{
		Cfg: &cfg,
	}

	app.Run(app.Mount())

}
