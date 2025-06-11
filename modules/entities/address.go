package entities

type Address struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	Type          string
	RecipientName string
	LastName      string
	Phone         string
	Province      string
	Address       string
	Other         string
	District      string
	SubDistrict   string
	PostalCode    int

	ProvinceId int
	AmphureId  int
	TambonId   int
	IsDefault  bool
}
