package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindOneUserByemail(email string) (*entities.User, error)
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) FindOneUserByemail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) UpdateUser(user *entities.User) error {
	return r.db.Save(user).Error
}
