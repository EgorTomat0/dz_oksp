package user

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	Id               int         `json:"id"`
	Name             string      `json:"uname"`
	RegistrationDate pgtype.Date `json:"reg_date"`
}
