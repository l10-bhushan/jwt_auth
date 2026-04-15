package handler

import (
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

// Todo: Implement login handler
func (userHandler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {}

// Todo: Implement Signup handler
func (userHandler *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {}
