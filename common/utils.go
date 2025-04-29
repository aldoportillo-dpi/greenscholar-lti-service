package common

import (
	"github.com/go-chi/chi/v5"
)

func GetAPIBaseRouter() *chi.Mux {
	apiRouter := chi.NewRouter()

	return apiRouter
}
