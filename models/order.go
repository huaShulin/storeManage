package models

type OrderAddGoodsParam struct {
	Id	string `form:"id"`
	Number int `form:"number"`
}

type OrderGoodsResult struct {
	OrderGoods		[]OrderGoods		`json:"rows"`
	Total		int			`json:"total"`
}
type OrderGoods struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Price 		float64		`json:"price"`
	OrderNumber	int			`json:"orderNumber"`
	Number		int			`json:"number"`
	Type		string		`json:"type"`
}