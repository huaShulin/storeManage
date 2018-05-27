package models

type OrderAddGoodsParam struct {
	Id	string `form:"id"`
	Number int `form:"number"`
}

type OrderRemarkParam struct {
	Remark string `form:"remark"`
}

type OrderResult struct {
	Order		[]Order		`json:"rows"`
	Total		int			`json:"total"`
}

type Order struct {
	Id			string		`json:"id"`
	Sum			float64		`json:"sum"`
	UserId		string		`json:"userId"`
	UserName	string		`json:"userName"`
	Time		string		`json:"time"`
	Remark		string		`json:"remark"`
	Details		string		`json:"details"`
}

type OrderGoodsResult struct {
	OrderGoods		[]OrderGoods		`json:"rows"`
	Total			int			`json:"total"`
}
type OrderGoods struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Price 		float64		`json:"price"`
	OrderNumber	int			`json:"orderNumber"`
	Number		int			`json:"number"`
	Type		string		`json:"type"`
}