package fascade

import (
	"context"

	"github.com/miguelmartinez624/web-service-seed/domains/autors"
	"github.com/miguelmartinez624/web-service-seed/domains/movies"
)

type Movies struct {
	movieService  *movies.Service
	autorsService *autors.Service
}

func NewMovies(movieStore movies.Store, autorsStore autors.Store) *Movies {
	//Servics initialization
	service := movies.NewService(movieStore)
	autorService := autors.NewService(autorsStore)

	Movies := Movies{movieService: service, autorsService: autorService}
	return &Movies
}

func (m *Movies) GetMovies(ctx context.Context) (movies []movies.Movie, err error) {
	return m.movieService.GetAll(ctx)
}

func (m *Movies) CreateMovie(ctx context.Context, movie *movies.Movie) (ID string, err error) {

	autor, err := m.autorsService.FindAutorByID(ctx, movie.AutorID)
	if err != nil {
		return "", err
	}

	//If the autor dosent Exist
	if autor == nil {
		return "", autors.AutorNoExistError
	}

	return m.movieService.CreateMovie(ctx, movie)
}
