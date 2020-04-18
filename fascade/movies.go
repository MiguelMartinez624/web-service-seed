package fascade

import (
	"context"

	"github.com/miguelmartinez624/web-service-seed/domains/movies"
)

type Movies struct {
	movieService *movies.Service
}

func NewMovies(movieStore movies.Store) *Movies {
	service := movies.NewService(movieStore)
	Movies := Movies{movieService: service}
	return &Movies
}

func (m *Movies) GetMovies(ctx context.Context) (movies []movies.Movie, err error) {
	return m.movieService.GetAll(ctx)
}

func (m *Movies) CreateMovie(ctx context.Context, movie *movies.Movie) (ID string, err error) {
	return m.movieService.CreateMovie(ctx, movie)
}
