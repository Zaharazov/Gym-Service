package models

import "database/sql"

type User struct {
	ID       int
	Login    string
	Password string
	Active   sql.NullBool
	Name     string
	Age      int
	Sex      string
	Phone    string
	PassID   sql.NullInt64
	GymID    sql.NullInt64
}
