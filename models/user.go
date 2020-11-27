package models

type User struct {
	ID       int
	Name     string
	Surname  string
	Age      int
	Gender   string
	Login    string
	Password string
	Remove   bool
}
