package models

type GoodsResult struct {
	Goods		[]Goods		`json:"goods"`
	Result
}

type Goods struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Price 		float64		`json:"price"`
	Remark 		string		`json:"remark"`
	TypeId		string		`json:"typeId"`
}