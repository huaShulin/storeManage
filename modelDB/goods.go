package modelDB

const (
	GET_GOODS_LIST = " SELECT G.ID AS ID, G.NAME AS NAME, G.PRICE AS PRICE, G.SRC AS SRC, G.REMARK AS REMARK, G.NUMBER AS NUMBER, G.PURCHASE AS PURCHASE, T.NAME AS TYPE FROM TB_GOODS G LEFT JOIN TB_TYPE T ON G.TYPE_ID = T.ID "
)

type Goods struct {
	Id 			string		`gorm:"column:ID" json:"ID"`
	Name 		string		`gorm:"column:NAME" json:"NAME"`
	Price 		float64		`gorm:"column:PRICE" json:"PRICE"`
	Src			string		`gorm:"column:SRC" json:"SRC"`
	Remark 		string		`gorm:"column:REMARK" json:"REMARK"`
	Number		int			`gorm:"column:NUMBER" json:"NUMBER"`
	Purchase	int			`gorm:"column:PURCHASE" json:"PURCHASE"`
	Type		string		`gorm:"column:TYPE" json:"TYPE"`
}

type SaveGoods struct {
	Id 			string		`gorm:"column:ID" json:"ID"`
	Name 		string		`gorm:"column:NAME" json:"NAME"`
	Price 		float64		`gorm:"column:PRICE" json:"PRICE"`
	Src			string		`gorm:"column:SRC" json:"SRC"`
	Remark 		string		`gorm:"column:REMARK" json:"REMARK"`
	Number		int			`gorm:"column:NUMBER" json:"NUMBER"`
	Purchase	int			`gorm:"column:PURCHASE" json:"PURCHASE"`
	TypeId		string		`gorm:"column:TYPE_ID" json:"TYPE_ID"`
}