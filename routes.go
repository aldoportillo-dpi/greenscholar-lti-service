package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/illinoisdpi/go-service-template/api/users"
	"github.com/illinoisdpi/go-service-template/api/wire"
	"github.com/illinoisdpi/go-service-template/common"
)

func initHTTPHandlers(r *chi.Mux) {
	// unauthenticated routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Running on port %s 🚀", os.Getenv("PORT"))))
	})

	// authenticated routes
	apiRouter := common.GetAPIBaseRouter()
	users.InitUserHandlers(apiRouter, wire.InitializeUserService())

	r.Mount("/api", apiRouter)
}
