package entities

type Category struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Icon string
}
