package models

import "database/sql"

type Coach struct {
	ID          int
	Login       string
	Password    string
	Name        string
	Age         int
	Sex         string
	Description string
	GymID       sql.NullInt64
}
