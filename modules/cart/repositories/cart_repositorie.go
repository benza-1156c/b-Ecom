package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type CartRepository interface {
	Create(data *entities.Cart) error
	CreateCartItem(data *entities.CartItem) error
	FindCartByUserId(id uint) (*entities.Cart, error)
	FindByCartIDAndProductID(userID, productID uint) (*entities.CartItem, error)
	FindAllCartItemByCartid(id uint) ([]entities.CartItem, error)
	UpdateCartItem(data *entities.CartItem) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) Create(data *entities.Cart) error {
	return r.db.Create(data).Error
}

func (r *cartRepository) CreateCartItem(data *entities.CartItem) error {
	return r.db.Create(data).Error
}

func (r *cartRepository) FindCartByUserId(id uint) (*entities.Cart, error) {
	cart := &entities.Cart{}
	if err := r.db.Where("user_id = ?", id).First(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *cartRepository) FindByCartIDAndProductID(cartID, productID uint) (*entities.CartItem, error) {
	cart := &entities.CartItem{}
	if err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *cartRepository) FindAllCartItemByCartid(id uint) ([]entities.CartItem, error) {
	var cart []entities.CartItem
	if err := r.db.Preload("Product.Images").Where("cart_id = ?", id).Find(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *cartRepository) UpdateCartItem(data *entities.CartItem) error {
	return r.db.Save(data).Error
}
