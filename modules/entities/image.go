package entities

type Image struct {
	ID        uint `gorm:"primaryKey"`
	Url       string
	ProductID uint
}
