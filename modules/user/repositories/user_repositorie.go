package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindOneById(id uint) (*entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindOneById(id uint) (*entities.User, error) {
	user := &entities.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
