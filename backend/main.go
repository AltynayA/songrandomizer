package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"songretriever/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "444"
	dbname   = "spotify-songs"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	dbclient := db.NewDbClient(psqlInfo, logger)
	defer dbclient.Shutdown()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/songs" , func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		songs, err:= dbclient.GetSongs() 
		if err != nil {
			w.Write([]byte(err.Error()))
			return 
		}
		x,err := json.Marshal(songs)
		if err != nil {
			w.Write([]byte(err.Error()))
			return 
		}
		w.Write(x)
	})

	http.ListenAndServe(":3000", r)
}

