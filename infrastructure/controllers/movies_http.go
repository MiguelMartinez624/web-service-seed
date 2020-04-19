package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/miguelmartinez624/web-service-seed/domains/movies"
	"github.com/miguelmartinez624/web-service-seed/fascade"
)

type MoviesHttp struct {
	moviesCases *fascade.Movies
}

func NewMoviesHttp(moviesCases *fascade.Movies) *MoviesHttp {
	return &MoviesHttp{moviesCases: moviesCases}
}

func (m MoviesHttp) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := m.moviesCases.GetMovies(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(movies)

}

func (m MoviesHttp) CreateHandler(w http.ResponseWriter, r *http.Request) {
	// Declare a new Movie struct.
	var movie movies.Movie

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ID, err := m.moviesCases.CreateMovie(r.Context(), &movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "New Movie ID: %+v", ID)
}
