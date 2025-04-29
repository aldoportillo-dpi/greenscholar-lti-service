package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var userService *UserService

// InitUserHandlers initializes the user handlers
func InitUserHandlers(r *chi.Mux, service *UserService) {
	userService = service

	// initialize http handlers
	r.Route("/v1/users", func(r chi.Router) {
		r.Get("/", listUserHandler)
	})
}

func listUserHandler(w http.ResponseWriter, r *http.Request) {

}
