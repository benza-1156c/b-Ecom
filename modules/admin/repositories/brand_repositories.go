package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type BrandRepository interface {
	Create(data *entities.Brand) error
	FindAll() ([]entities.Brand, error)
	Delete(id uint) error
	Update(id uint, data *entities.Brand) error
}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{db}
}

func (r *brandRepository) Create(data *entities.Brand) error {
	return r.db.Create(data).Error
}

func (r *brandRepository) FindAll() ([]entities.Brand, error) {
	var brand []entities.Brand
	if err := r.db.Find(&brand).Error; err != nil {
		return nil, err
	}
	return brand, nil
}

func (r *brandRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Brand{}, id).Error
}

func (r *brandRepository) Update(id uint, data *entities.Brand) error {
	brand := &entities.Brand{}
	if err := r.db.First(brand, id).Error; err != nil {
		return err
	}

	brand.Name = data.Name
	brand.Icon = data.Icon

	return r.db.Save(brand).Error
}
