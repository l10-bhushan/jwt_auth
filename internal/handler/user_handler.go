package handler

import (
	"encoding/json"
	"net/http"

	"github.com/l10-bhushan/jwt_auth/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (usehandler *UserHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "Server is healthy",
	})
}

// Todo: Implement login handler
func (userHandler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {}

// Todo: Implement Signup handler
func (userHandler *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {}
