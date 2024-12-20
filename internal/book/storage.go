package book

import "context"

type Repo interface {
	Create(ctx context.Context, book Book) (string, error)
	GetBook(ctx context.Context, id string) (Book, error)
	Update(ctx context.Context, book Book) error
	Delete(ctx context.Context, id string) error
}
