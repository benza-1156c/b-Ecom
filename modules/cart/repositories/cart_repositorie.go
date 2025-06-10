package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type CartRepository interface {
	FnidUserById(id uint) (*entities.User, error)
	Create(data *entities.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) FnidUserById(id uint) (*entities.User, error) {
	user := &entities.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *cartRepository) Create(data *entities.Cart) error {
	return r.db.Create(data).Error
}
