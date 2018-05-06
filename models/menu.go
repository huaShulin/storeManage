package models

type MenuResult struct {
	Menus		[]Menu		`json:"menus"`
	Result
}

type Menu struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	IconCls 	string 		`json:"iconCls"`
	Href 		string		`json:"href"`
	Children 	[]Menu		`json:"children"`
}