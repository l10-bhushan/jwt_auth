package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Server could'nt start...")
		}
	}()

	// Listen for OS signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Block until signal received
	log.Println("Shutting down server...")

	// Create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	} else {
		log.Println("server stopped gracefully")
	}

}
