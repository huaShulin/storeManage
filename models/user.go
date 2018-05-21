package models

type Login struct {
	Phone string `json:"phone"`
	Password string	`json:"password"`
}

type UserResult struct {
	Users		[]User		`json:"rows"`
	Total		int			`json:"total"`
}

type User struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Phone		string		`json:"phone"`
	Status		string		`json:"status"`
	Role 		string		`json:"role"`
}

type UserParam struct {
	Name 		string		`form:"name"`
	Phone 		string		`form:"phone"`
	Status		int			`form:"status"`
	RoleIds		[]string	`form:"roleIds"`
}