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