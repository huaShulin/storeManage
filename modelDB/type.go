package modelDB

type Type struct {
	Id 			string		`gorm:"column:ID" json:"ID"`
	Name 		string		`gorm:"column:NAME" json:"NAME"`
	ParentId 	string 		`gorm:"column:PARENT_ID" json:"PARENT_ID"`
}