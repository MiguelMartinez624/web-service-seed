package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/web-service-seed/fascade"
	"github.com/miguelmartinez624/web-service-seed/infrastructure/controllers"
	"github.com/miguelmartinez624/web-service-seed/infrastructure/persistency"
)

const DB_URI = "mongodb://localhost:27017"

func main() {

	client, cancel := ConnectMongoDB(DB_URI)
	defer cancel()

	moviesStore := persistency.NewMongoDBMovieStore(client.Database("kuebana"))
	module := fascade.NewMovies(moviesStore)
	httpController := controllers.NewMoviesHttp(module)

	//we add endpoints here to mux

	r := mux.NewRouter()
	r.HandleFunc("/movies", httpController.GetMoviesHandler).Methods("GET")
	r.HandleFunc("/movies", httpController.CreateHandler).Methods("POST")

	http.Handle("/", r)

}
