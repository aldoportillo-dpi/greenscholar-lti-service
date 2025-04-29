package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/illinoisdpi/go-service-template/configs"
	"github.com/illinoisdpi/go-service-template/db"
)

func main() {
	// initialize logger
	logger := configs.NewLogger()

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Error(".env file not found", "err", err)
	}

	// run migrations
	logger.Info("Running migrations")
	err = db.RunMigrations()
	if err != nil {
		log.Fatal(err)
	}

	// initialize router
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.AllowAll().Handler)

	// handlers
	initHTTPHandlers(r)

	logger.Info("Starting server", "port", os.Getenv("PORT"))

	// start server
	err = http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		panic(err)
	}
}
