package entities

import "time"

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Status    string
	Total     int
	CreatedAt time.Time

	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     int

	Product Product `gorm:"foreignKey:ProductID"`
}
