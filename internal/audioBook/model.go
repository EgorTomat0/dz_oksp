package audioBook

import "github.com/jackc/pgx/v5/pgtype"

type AudioBook struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
	ReleaseDate pgtype.Date `json:"release_date"`
	Rating      float32     `json:"rating"`
	CoverImage  string      `json:"cover_image"`
	AudioFile   string      `json:"audio_file"`
}
