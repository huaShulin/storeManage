package models

type Result struct {
	Success bool		`json:"success"`
	Message string		`json:"message"`
}

type IdParam struct {
	Id	string `form:"id"`
}

type PageInfo struct {
	Page int `form:"page"`
	Rows int `form:"rows"`
}
