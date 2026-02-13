package routes

import (
	"backend_bench/internal/handler/health"
	"net/http"
)

func RegisterUsersRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", health.Handler)
}
