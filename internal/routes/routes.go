package routes

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	RegisterUsersRoutes(mux)
	RegisterStatusRoutes(mux)
}
