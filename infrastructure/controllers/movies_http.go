package controllers

import (
	"net/http"

	"github.com/miguelmartinez624/web-service-seed/fascade"
)

type MoviesHttp struct {
	moviesCases *fascade.Movies
}

func NewMoviesHttp(moviesCases *fascade.Movies) *MoviesHttp {
	return &MoviesHttp{moviesCases: moviesCases}
}

func (m MoviesHttp) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {

}
func (m MoviesHttp) CreateHandler(w http.ResponseWriter, r *http.Request) {

}
