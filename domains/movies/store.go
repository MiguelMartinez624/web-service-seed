package movies

import "context"

type Store interface {
	GetAll(ctx context.Context) (movie []Movie, err error)
	GetByID(ctx context.Context, ID string) (movie *Movie, err error)
	Save(ctx context.Context, movie *Movie) (ID string, err error)
}
