package entities

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       int
	Count       int
	Status      string
	Sku         string
	Featured    bool

	Images []Image `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`

	CategoriesID uint
	Categories   Category `gorm:"foreignKey:CategoriesID"`

	BrandID uint
	Brand   Brand `gorm:"foreignKey:BrandID"`

	OrderItems []OrderItem `gorm:"foreignKey:ProductID" json:"-"`
	CartItems  []CartItem  `gorm:"foreignKey:ProductID" json:"-"`
}
