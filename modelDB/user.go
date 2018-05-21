package modelDB

const (
	//GET_USER_LIST = " SELECT U.ID AS ID, U.NAME AS NAME, U.PHONE AS PHONE, U.PASSWORD AS PASSWORD, U.STATUS AS STATUS, R.NAME AS ROLE FROM TB_USER U LEFT JOIN TB_ROLE R ON U.ROLE_ID = R.ID "
	GET_USER_LIST = " SELECT ID, NAME, PHONE, PASSWORD, STATUS FROM TB_USER "
)

type User struct {

	// ID
	Id string `gorm:"column:ID" json:"ID"`

	// 用户名
	Name string `gorm:"column:NAME" json:"NAME"`

	// 手机号
	Phone string `gorm:"column:PHONE" json:"PHONE"`

	// 密码
	Password string `gorm:"column:PASSWORD" json:"PASSWORD"`

	//状态
	Status int `gorm:"column:STATUS" json:"STATUS"`

	// 角色
	RoleID string `gorm:"column:ROLE_ID" json:"ROLE_ID"`

}

type UserList struct {

	// ID
	Id string `gorm:"column:ID" json:"ID"`

	// 用户名
	Name string `gorm:"column:NAME" json:"NAME"`

	// 手机号
	Phone string `gorm:"column:PHONE" json:"PHONE"`

	// 密码
	Password string `gorm:"column:PASSWORD" json:"PASSWORD"`

	//状态
	Status int `gorm:"column:STATUS" json:"STATUS"`

	// 角色
	Role string `gorm:"column:ROLE" json:"ROLE"`

}

type SaveUser struct {
	Id 			string		`gorm:"column:ID" json:"ID"`
	Name 		string		`gorm:"column:NAME" json:"NAME"`
	Price 		float64		`gorm:"column:PRICE" json:"PRICE"`
	Src			string		`gorm:"column:SRC" json:"SRC"`
	Remark 		string		`gorm:"column:REMARK" json:"REMARK"`
	Number		int			`gorm:"column:NUMBER" json:"NUMBER"`
	Purchase	int			`gorm:"column:PURCHASE" json:"PURCHASE"`
	TypeId		string		`gorm:"column:TYPE_ID" json:"TYPE_ID"`
}