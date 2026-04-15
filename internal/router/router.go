package router

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/l10-bhushan/jwt_auth/internal/handler"
	"github.com/l10-bhushan/jwt_auth/internal/repository"
	"github.com/l10-bhushan/jwt_auth/internal/service"
)

type Config struct {
	Addr string
}

type Application struct {
	Cfg *Config
}

func NewAppliction(cfg *Config) *Application {
	return &Application{
		Cfg: cfg,
	}
}

func (app *Application) Mount() http.Handler {
	router := chi.NewRouter()

	repo := repository.NewInMemoryUserRepo()
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	router.Get("/health", handler.Health)

	return router
}

func (app *Application) Run(router http.Handler) {
	server := &http.Server{
		Addr:         app.Cfg.Addr,
		Handler:      router,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server could'nt start...")
	}
}
