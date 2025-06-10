package reqaddress

type Reqaddress struct {
	Type          string `json:"type"`
	RecipientName string `json:"recipientName"`
	LastName      string `json:"lastName"`
	Phone         string `json:"phone"`
	Province      string `json:"province"`
	District      string `json:"district"`
	SubDistrict   string `json:"subDistrict"`
	Other         string `json:"other"`
	PostalCode    int    `json:"postalCode"`
	IsDefault     bool   `json:"isDefault"`
}
