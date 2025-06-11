package reqaddress

type Reqaddress struct {
	Type          string `json:"type"`
	RecipientName string `json:"recipientName"`
	LastName      string `json:"lastName"`
	Phone         string `json:"phone"`
	Province      string `json:"province"`
	District      string `json:"district"`
	Address       string `json:"address"`
	SubDistrict   string `json:"subDistrict"`
	Other         string `json:"other"`
	PostalCode    int    `json:"postalCode"`
	IsDefault     bool   `json:"isDefault"`
	ProvinceId    int    `json:"provinceId"`
	AmphureId     int    `json:"amphureId"`
	TambonId      int    `json:"tambonId"`
}
