package req

type ReqCategory struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type ReqBrand struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type ReqProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Sku         string `json:"sku"`
	Price       int    `json:"price"`
	Count       int
	Images      []string `json:"images"`
	Category    uint     `json:"category"`
	Brand       uint     `json:"brand"`
	Featured    bool     `json:"featured"`
	Status      string
}

type ReqUser struct {
	Status string `json:"status"`
}
