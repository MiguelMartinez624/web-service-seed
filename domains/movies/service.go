package movies

import (
	"context"
)

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}
func (s *Service) GetAll(ctx context.Context) (movies []Movie, err error) {
	movies, err = s.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return movies, err
}

func (s *Service) CreateMovie(ctx context.Context, movie *Movie) (ID string, err error) {
	ID, err = s.store.Save(ctx, movie)
	return ID, err
}
