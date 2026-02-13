package routes

import (
	"backend_bench/internal/handler/status"
	"net/http"
)

func RegisterStatusRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/status", status.StatusHandler)
}
