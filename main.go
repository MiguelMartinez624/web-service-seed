package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/miguelmartinez624/web-service-seed/fascade"
	"github.com/miguelmartinez624/web-service-seed/infrastructure/controllers"
	"github.com/miguelmartinez624/web-service-seed/infrastructure/persistency"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Up and running")
	log.Fatal(srv.ListenAndServe())

}

func ConnectMongoDB(uri string) (client *mongo.Client, cancel context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, cancel
}
