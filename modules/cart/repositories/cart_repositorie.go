package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type CartRepository interface {
	Create(data *entities.Cart) error
	CreateCartItem(data *entities.CartItem) error
	FindCartByUserId(id uint) (*entities.Cart, error)
	FindCartItemById(id uint) (*entities.CartItem, error)
	FindByCartIDAndProductID(userID, productID uint) (*entities.CartItem, error)
	FindAllCartItemByCartid(id uint) ([]entities.CartItem, error)
	UpdateCartItem(data *entities.CartItem) error
	UpdateCountCartItem(data *entities.CartItem) error
	DeleteCartItemByID(id uint) error
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

func (r *cartRepository) UpdateCountCartItem(data *entities.CartItem) error {
	return r.db.Model(&entities.CartItem{}).
		Where("id = ?", data.ID).
		Update("quantity", data.Quantity).Error
}

func (r *cartRepository) FindCartItemById(id uint) (*entities.CartItem, error) {
	data := &entities.CartItem{}
	if err := r.db.Preload("Cart").First(data, id).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *cartRepository) DeleteCartItemByID(id uint) error {
	return r.db.Delete(&entities.CartItem{}, id).Error
}
