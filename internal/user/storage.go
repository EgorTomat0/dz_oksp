package user

import "context"

type Repo interface {
	Create(ctx context.Context, user User) (string, error)
	GetUser(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
