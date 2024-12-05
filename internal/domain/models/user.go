package models

type User struct {
	ID       int
	Login    string
	Password string
	Active   bool
	Name     string
	Age      int
	Sex      string
	Phone    string
	PassID   int
	GymID    int
}
