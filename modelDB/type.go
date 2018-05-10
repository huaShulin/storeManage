package modelDB

const (
	GET_TYPE_LIST = " SELECT C.ID AS ID, C.NAME AS NAME, P.NAME AS PARENT FROM TB_TYPE C LEFT JOIN TB_TYPE P ON C.PARENT_ID = P.ID "
)

type Type struct {
	Id 			string		`gorm:"column:ID" json:"ID"`
	Name 		string		`gorm:"column:NAME" json:"NAME"`
	ParentId 	string 		`gorm:"column:PARENT_ID" json:"PARENT_ID"`
}

type TypePage struct {
	Id 			string		`gorm:"column:ID" json:"ID"`
	Name 		string		`gorm:"column:NAME" json:"NAME"`
	Parent	 	string 		`gorm:"column:PARENT" json:"PARENT"`
}