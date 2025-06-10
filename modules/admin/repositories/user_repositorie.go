package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(role, search string, page, limit int) ([]entities.User, int64, error)
	Delete(id uint) error
	UpdateStatus(id uint, user *entities.User) error
	FindTotal() (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll(role, search string, page, limit int) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	query := r.db.Model(&entities.User{})

	if role != "" {
		query = query.Where("role = ?", role)
	}

	if search != "" {
		query = query.Where("username ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) Delete(id uint) error {
	if err := r.db.Delete(&entities.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateStatus(id uint, user *entities.User) error {
	existingUser := &entities.User{}
	if err := r.db.First(existingUser, id).Error; err != nil {
		return err
	}

	existingUser.Status = user.Status

	if err := r.db.Save(existingUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindTotal() (int64, error) {
	var total int64
	if err := r.db.Model(&entities.User{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
