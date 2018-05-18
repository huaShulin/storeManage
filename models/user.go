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
	Price 		float64		`form:"price"`
	Remark 		string		`form:"remark"`
	Number		int			`form:"number"`
	Purchase	int			`form:"purchase"`
	TypeId		string		`form:"typeId"`
}