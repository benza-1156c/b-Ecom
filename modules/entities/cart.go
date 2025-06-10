package entities

import "time"

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	CreatedAt time.Time

	CartItems []CartItem `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE;"`
}

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint
	ProductID uint
	Quantity  int

	Product Product `gorm:"foreignKey:ProductID"`
}
