package user

import (
	"context"
	"dz_oksp/pkg/pgsql"
	"strconv"
)

type userRepo struct {
	client pgsql.Conn
}

func (u *userRepo) Create(ctx context.Context, user User) (string, error) {
	q := `INSERT INTO public.users (name, reg_date) VALUES ($1, CURRENT_DATE) RETURNING id`
	if err := u.client.QueryRow(ctx, q, user.Name).Scan(&user.Id); err != nil {
		return "", err
	}
	return strconv.Itoa(user.Id), nil
}

func (u *userRepo) GetUser(ctx context.Context, id string) (User, error) {
	q := `SELECT * FROM public.users WHERE id = $1`
	rows, err := u.client.Query(ctx, q, id)
	defer rows.Close()
	if err != nil {
		return User{}, err
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.RegistrationDate)
		if err != nil {
			return User{}, err
		}
	}
	return user, err
}

func (u *userRepo) Update(ctx context.Context, user User) error {
	q := `UPDATE public.users SET name = $1 WHERE id = $2`
	_, err := u.client.Exec(ctx, q, user.Name, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM public.users WHERE id = $1`
	_, err := u.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

func NewRepo(client pgsql.Conn) Repo {
	return &userRepo{client: client}
}
