package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"songretriever/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func main() {
	// Try to read from environment variables first (Docker)
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// Fallback to local values (for running without Docker)
	if host == "" {
		host = "localhost"
		port = "5433"
		user = "postgres"
		password = "444"
		dbname = "spotify-songs"
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	dbclient := db.NewDbClient(psqlInfo, logger)
	if dbclient == nil {
		logger.Fatal("Failed to connect to database")
	}
	defer dbclient.Shutdown()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/songs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		songs, err := dbclient.GetSongs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		x, err := json.Marshal(songs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(x)
	})

	http.ListenAndServe(":3000", r)
}
