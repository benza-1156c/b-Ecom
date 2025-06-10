package entities

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex"`
	UserName  string
	Avatar    string
	Role      string
	Status    string
	CreatedAt time.Time

	Orders  []Order   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Cart    Cart      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Address []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
