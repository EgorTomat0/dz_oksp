package book

import "github.com/jackc/pgx/v5/pgtype"

type Book struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
	ReleaseDate pgtype.Date `json:"release_date"`
	Rating      float32     `json:"rating"`
	CoverImage  string      `json:"cover_image"`
	Paragraph   pgtype.Text `json:"paragraph"`
}
