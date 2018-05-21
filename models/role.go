package models

type RoleResult struct {
	Roles		[]Role		`json:"rows"`
	Total		int			`json:"total"`
}

type Role struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Menu		string		`json:"menu"`
	Number		int			`json:"number"`
}

type RoleParam struct {
	Name 		string		`form:"name"`
	MenuIds		[]string	`form:"menuIds"`
}