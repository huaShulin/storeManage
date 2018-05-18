package modelDB

const (
	GET_MENU_LIST = " SELECT ID, NAME FROM TB_ROLE "
)

type Menu struct {

	// ID
	Id string `gorm:"column:ID" json:"ID"`

	// 目录名
	Name string `gorm:"column:NAME" json:"NAME"`

	// 图标
	IconCls string `gorm:"column:ICON_CLS" json:"ICON_CLS"`

	// 连接
	Href string `gorm:"column:HREF" json:"HREF"`

	// 父ID
	ParentId string `gorm:"column:PARENT_ID" json:"PARENT_ID"`
}