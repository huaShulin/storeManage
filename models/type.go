package models

type TypeResult struct {
	Type		[]Type		`json:"menus"`
	Result
}

type Type struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`

	ParentId 	string 		`json:"parentId"`
}

type TypePageResult struct {
	TypePages	[]TypePage	`json:"rows"`
	Total		int			`json:"total"`
}

type TypePage struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Number		int		`json:"number"`
	Parent 		string 		`json:"parent"`
}

type TypeParam struct {
	Name 		string		`form:"name"`
	ParentId	string		`form:"parentId"`
}