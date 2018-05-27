package modelDB

const (
	GET_ORDER_LIST = " SELECT ID,SUM,USER_ID,USER_NAME,TIME,REMARK FROM TB_ORDER "
)

type Order struct {
	Id			string		`gorm:"column:ID" json:"ID"`
	Sum			float64		`gorm:"column:SUM" json:"SUM"`
	UserId		string		`gorm:"column:USER_ID" json:"USER_ID"`
	UserName	string		`gorm:"column:USER_NAME" json:"USER_NAME"`
	Time		string		`gorm:"column:TIME" json:"TIME"`
	Remark		string		`gorm:"column:REMARK" json:"REMARK"`
}

type OrderGoods struct {
	Id 			string		`gorm:"column:ID" json:"ID"`
	Name 		string		`gorm:"column:NAME" json:"NAME"`
	Price 		float64		`gorm:"column:PRICE" json:"PRICE"`
	Number		int			`gorm:"column:NUMBER" json:"NUMBER"`
	Type		string		`gorm:"column:TYPE" json:"TYPE"`
	Sum			float64		`gorm:"column:SUM" json:"SUM"`
	GoodsId		string		`gorm:"column:GOODS_ID" json:"GOODS_ID"`
	OrderId     string		`gorm:"column:ORDER_ID" json:"ORDER_ID"`
}