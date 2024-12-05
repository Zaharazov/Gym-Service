package models

type Admin struct {
	ID          int
	Login       string
	Password    string
	Name        string
	AccessLevel int
	Phone       string
	GymID       int
}
