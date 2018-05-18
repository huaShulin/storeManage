package modelDB

const (
	GET_ROLE_LIST = " SELECT ID, NAME FROM TB_ROLE "
)

type RoleList struct {

	// ID
	Id string `gorm:"column:ID" json:"ID"`

	// 职位
	Name string `gorm:"column:NAME" json:"NAME"`

}

type SaveRole struct {
	Id 			string		`gorm:"column:ID" json:"ID"`
	Name 		string		`gorm:"column:NAME" json:"NAME"`
}