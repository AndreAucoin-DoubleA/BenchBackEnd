package routes

import (
	"backend_bench/internal/handler/status"
	"backend_bench/internal/middleware"

	"net/http"
)

func RegisterStatusRoutes(mux *http.ServeMux, jwtSecret string) {
	mux.Handle("/status", middleware.AuthMiddleware(http.HandlerFunc(status.StatusHandler), jwtSecret))
}
