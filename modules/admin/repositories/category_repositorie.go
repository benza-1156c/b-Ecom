package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(data *entities.Category) error
	Update(id uint, data *entities.Category) error
	FindAll() ([]entities.Category, error)
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(data *entities.Category) error {
	return r.db.Create(data).Error
}

func (r *categoryRepository) Update(id uint, data *entities.Category) error {
	var existing entities.Category
	if err := r.db.First(&existing, id).Error; err != nil {
		return err
	}

	existing.Name = data.Name
	existing.Icon = data.Icon

	return r.db.Save(&existing).Error
}

func (r *categoryRepository) FindAll() ([]entities.Category, error) {
	var category []entities.Category
	if err := r.db.Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Category{}, id).Error
}
