package reqcart

type ReqCart struct {
	ProductID uint `json:"productid"`
	Quantity  int  `json:"quantity"`
}

type ReqUpdateCart struct {
	CartId    uint `json:"cartId"`
	Quantity  int  `json:"quantity"`
	ProductId uint `json:"productId"`
}
