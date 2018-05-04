package models

type Login struct {
	Username string
	Password string
}

type Result struct {
	Success bool		`json:"success"`
	Message string		`json:"message"`
}
