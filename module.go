package main

import (
	"context"

	"github.com/miguelmartinez624/web-service-seed/domains/movies"
)

type Module struct {
	movieService movies.Service
}

func (m *Module) GetMovies(ctx context.Context) (movies []movies.Movie, err error) {
	return
}
