package entities

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       int
	Count       int
	Featured    bool
	Status      string
	Sku         string

	Images []Image `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`

	CategoriesID uint
	Categories   Category `gorm:"foreignKey:CategoriesID"`

	BrandID uint
	Brand   Brand `gorm:"foreignKey:BrandID"`

	OrderItems []OrderItem `gorm:"foreignKey:ProductID"`
	CartItems  []CartItem  `gorm:"foreignKey:ProductID"`
}
