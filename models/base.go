package models

type Result struct {
	Success bool		`json:"success"`
	Message string		`json:"message"`
}

type PageInfo struct {
	Page int `form:"page"`
	Rows int `form:"rows"`
}
