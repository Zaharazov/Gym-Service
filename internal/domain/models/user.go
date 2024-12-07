package models

import "database/sql"

type User struct {
	ID       int
	Login    string
	Password string
	Active   sql.NullBool
	Name     string
	Age      sql.NullInt64
	Sex      sql.NullString
	Phone    sql.NullString
	PassID   sql.NullInt64
	GymID    sql.NullInt64
}
