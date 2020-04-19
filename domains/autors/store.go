package autors

import "context"

type Store interface {
	GetAll(ctx context.Context) (autor []Autor, err error)
	GetByID(ctx context.Context, ID string) (autor *Autor, err error)
	Save(ctx context.Context, autor *Autor) (ID string, err error)
}
