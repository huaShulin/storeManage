package modelDB

const (
	PAGE = " ORDER BY ID LIMIT ?,? "
)

type ResultNumber struct {
	Number 			int		`gorm:"column:NUMBER" json:"NUMBER"`
}

type ResultName struct {
	Name			string	`gorm:"column:NAME" json:"NAME"`
}

type ResultIds struct {
	UserId			string	`gorm:"column:USER_ID" json:"USER_ID"`
	RoleId			string	`gorm:"column:ROLE_ID" json:"ROLE_ID"`
	MenuId			string	`gorm:"column:MENU_ID" json:"MENU_ID"`
}