package book

import (
	"context"
	"dz_oksp/pkg/pgsql"
	"strconv"
)

type bookRepo struct {
	client pgsql.Conn
}

func (b *bookRepo) Create(ctx context.Context, book Book) (string, error) {
	q := `INSERT INTO public.book (title, description, release_date, rating, cover, paragraph) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	if err := b.client.QueryRow(ctx, q, book.Title, book.Description, book.ReleaseDate, book.Rating, book.CoverImage, book.Paragraph).Scan(&book.Id); err != nil {
		return "", err
	}
	return strconv.Itoa(book.Id), nil
}

func (b *bookRepo) GetBook(ctx context.Context, id string) (Book, error) {
	q := `SELECT * FROM public.book WHERE id = $1`
	rows, err := b.client.Query(ctx, q, id)
	defer rows.Close()
	if err != nil {
		return Book{}, err
	}

	var book Book
	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.ReleaseDate, &book.Rating, &book.CoverImage, &book.Paragraph)
		if err != nil {
			return Book{}, err
		}
	}
	return book, err
}

func (b *bookRepo) Update(ctx context.Context, book Book) error {
	q := `UPDATE public.book SET title = $1, description = $2, release_date = $3, rating = $4, cover = $5, paragraph = $6 WHERE id = $7`
	_, err := b.client.Exec(ctx, q, book.Title, book.Description, book.ReleaseDate, book.Rating, book.CoverImage, book.Paragraph, book.Id)
	if err != nil {
		return err
	}
	return nil
}

func (b *bookRepo) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM public.book WHERE id = $1`
	_, err := b.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

func NewRepo(client pgsql.Conn) Repo {
	return &bookRepo{client: client}
}
