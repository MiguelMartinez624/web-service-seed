package autors

import "context"

type Service struct {
	store     Store
	validator *Validator
}

func NewService(store Store) *Service {
	return &Service{store: store, validator: &Validator{}}
}

func (s *Service) GetAll(ctx context.Context) (autors []Autor, err error) {
	autors, err = s.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return autors, err
}

func (s *Service) CreateAutor(ctx context.Context, autor *Autor) (ID string, err error) {

	err = s.validator.Validate(autor)
	if err != nil {
		return ID, err
	}

	ID, err = s.store.Save(ctx, autor)
	return ID, err
}
