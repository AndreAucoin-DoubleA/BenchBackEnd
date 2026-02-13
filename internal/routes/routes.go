package routes

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	RegisterStatusRoutes(mux)
}
