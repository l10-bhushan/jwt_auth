package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/l10-bhushan/jwt_auth/internal/auth"
)

type contextKey string

const UserContextKey contextKey = contextKey("user")

func JWTMiddleWare(next http.Handler) http.Handler {
	jwtService := auth.JWTService{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer")

		claims, err := jwtService.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
