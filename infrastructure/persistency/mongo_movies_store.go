package persistency

import (
	"context"

	"github.com/miguelmartinez624/web-service-seed/domains/movies"
)

type MongoDBMoviesStore struct{}

func (s *MongoDBMoviesStore) GetByID(ctx context.Context, ID string) (movie *movies.Movie, err error) {
	return
}
func (s *MongoDBMoviesStore) Save(movie *movies.Movie) (ID string, err error) {
	return
}
