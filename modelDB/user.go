package modelDB

type User struct {

	// ID
	Id string `gorm:"column:ID" json:"ID"`

	// 用户名
	Username string `gorm:"column:USERNAME" json:"USERNAME"`

	// 密码
	Password string `gorm:"column:PASSWORD" json:"PASSWORD"`

	// 角色ID
	RoleId string `gorm:"column:ROLE_ID" json:"ROLE_ID"`

}