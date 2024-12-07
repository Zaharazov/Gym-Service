package models

import "database/sql"

type Admin struct {
	ID          int
	Login       string
	Password    string
	Name        string
	AccessLevel sql.NullInt64
	Phone       sql.NullString
	GymID       sql.NullInt64
}
