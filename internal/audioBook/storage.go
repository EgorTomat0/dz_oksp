package audioBook

import "context"

type Repo interface {
	Create(ctx context.Context, ab AudioBook) (string, error)
	GetAB(ctx context.Context, id string) (AudioBook, error)
	Update(ctx context.Context, ab AudioBook) error
	Delete(ctx context.Context, id string) error
}
