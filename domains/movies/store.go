package movies

import "context"

type Store interface {
	GetByID(ctx context.Context, ID string) (movie *Movie, err error)
	Save(movie *Movie) (ID string, err error)
}
