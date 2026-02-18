package routes

import (
	"backend_bench/internal/model"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, repo *model.UserRepository, jwtSecret string) {
	RegisterStatusRoutes(mux, jwtSecret)
	RegisterLoginRoutes(mux, repo, jwtSecret)
}
