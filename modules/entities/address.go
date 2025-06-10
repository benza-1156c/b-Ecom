package entities

type Address struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	Type          string
	RecipientName string
	LastName      string
	Phone         string
	Province      string
	District      string
	SubDistrict   string
	PostalCode    int
	IsDefault     bool
}
