package models

import "database/sql"

type Admin struct {
	ID          int
	Login       string
	Password    string
	Name        string
	AccessLevel int
	Phone       string
	GymID       sql.NullInt64
}
