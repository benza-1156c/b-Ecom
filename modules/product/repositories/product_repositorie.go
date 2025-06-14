package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(search string, page, limit, categoryID, brandID, minPrice, maxPrice int) ([]entities.Product, int64, error)
	FindOneById(id uint) (*entities.Product, error)
	FindAllByCategory(id, Limit uint) ([]entities.Product, error)
	FindProductFeatured() ([]entities.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll(search string, page, limit, categoryID, brandID, minPrice, maxPrice int) ([]entities.Product, int64, error) {
	var products []entities.Product
	var total int64
	query := r.db.Model(&entities.Product{}).Preload("Images").Preload("Categories").Preload("Brand")

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if categoryID > 0 {
		query = query.Where("categories_id = ?", categoryID)
	}

	if brandID > 0 {
		query = query.Where("brand_id = ?", brandID)
	}

	if minPrice > 0 {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice > 0 {
		query = query.Where("price <= ?", maxPrice)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (r *productRepository) FindOneById(id uint) (*entities.Product, error) {
	product := &entities.Product{}
	if err := r.db.Preload("Images").First(product, id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) FindAllByCategory(id, limit uint) ([]entities.Product, error) {
	var products []entities.Product
	if err := r.db.
		Where("categories_id = ?", id).
		Limit(int(limit)).
		Preload("Images").
		Preload("Brand").
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) FindProductFeatured() ([]entities.Product, error) {
	var products []entities.Product
	if err := r.db.Preload("Images").Preload("Categories").Preload("Brand").Where("featured = ?", true).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
