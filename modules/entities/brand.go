package entities

type Brand struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Icon string
}
