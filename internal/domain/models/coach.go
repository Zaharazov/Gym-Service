package models

import "database/sql"

type Coach struct {
	ID          int
	Login       string
	Password    string
	Name        string
	Age         sql.NullInt64
	Sex         sql.NullString
	Description sql.NullString
	GymID       sql.NullInt64
}
