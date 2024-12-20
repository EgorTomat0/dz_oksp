package audioBook

import (
	"context"
	"dz_oksp/pkg/pgsql"
	"strconv"
)

type aBookRepo struct {
	client pgsql.Conn
}

func (a *aBookRepo) Create(ctx context.Context, ab AudioBook) (string, error) {
	q := `INSERT INTO public.book (title, description, release_date, rating, cover, paragraph) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	if err := a.client.QueryRow(ctx, q, ab.Title, ab.Description, ab.ReleaseDate, ab.Rating, ab.CoverImage, ab.AudioFile).Scan(&ab.Id); err != nil {
		return "", err
	}
	return strconv.Itoa(ab.Id), nil
}

func (a *aBookRepo) GetAB(ctx context.Context, id string) (AudioBook, error) {
	q := `SELECT * FROM public.book WHERE id = $1`
	rows, err := a.client.Query(ctx, q, id)
	defer rows.Close()
	if err != nil {
		return AudioBook{}, err
	}

	var aBook AudioBook
	for rows.Next() {
		err = rows.Scan(&aBook.Id, &aBook.Title, &aBook.Description, &aBook.ReleaseDate, &aBook.Rating, &aBook.CoverImage, &aBook.AudioFile)
		if err != nil {
			return AudioBook{}, err
		}
	}
	return aBook, err
}

func (a *aBookRepo) Update(ctx context.Context, ab AudioBook) error {
	q := `UPDATE public.book SET title = $1, description = $2, release_date = $3, rating = $4, cover = $5, audio = $6 WHERE id = $7`
	_, err := a.client.Exec(ctx, q, ab.Title, ab.Description, ab.ReleaseDate, ab.Rating, ab.CoverImage, ab.AudioFile, ab.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a *aBookRepo) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM public.abook WHERE id = $1`
	_, err := a.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

func NewRepo(client pgsql.Conn) Repo {
	return &aBookRepo{client: client}
}
