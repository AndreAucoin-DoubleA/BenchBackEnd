package middleware

import (
	"backend_bench/internal/auth"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		claims, err := auth.VerifyJWT(parts[1], secret)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Optionally store claims in context for handlers
		ctx := r.Context()
		ctx = context.WithValue(ctx, "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
