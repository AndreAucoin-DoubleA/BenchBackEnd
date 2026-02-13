package routes

import (
	"backend_bench/internal/handler/health"
	"net/http"
)

func RegisterStatusRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/status", health.Handler)
}
