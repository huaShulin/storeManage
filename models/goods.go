package models

type GoodsResult struct {
	Goods		[]Goods		`json:"goods"`
	Result
}

type Goods struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Price 		float64		`json:"price"`
	Src			string		`json:"src"`
	Remark 		string		`json:"remark"`
	Number		int			`json:"number"`
	Purchase	int			`json:"purchase"`
	Type		string		`json:"type"`
}

type SaveEditId struct {
	Id	string `json:"id"`
}