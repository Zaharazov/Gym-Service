package models

import "database/sql"

type Admin struct {
	ID       int
	Login    string
	Password string
	Name     string
	Phone    string
	GymID    sql.NullInt64
}
