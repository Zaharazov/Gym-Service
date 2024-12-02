package models

type Coach struct {
	ID          int
	Login       string
	Password    string
	Name        string
	Age         int
	Sex         string
	Description string
	GymID       int
}
