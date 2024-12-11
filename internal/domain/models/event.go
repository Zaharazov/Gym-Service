package models

import "database/sql"

type Event struct {
	ID          int
	Name        string
	Description string
	CoachID     sql.NullInt64
}
