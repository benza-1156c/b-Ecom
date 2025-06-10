package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(search string, page, limit, categoryID, brandID, price int) ([]entities.Product, int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll(search string, page, limit, categoryID, brandID, price int) ([]entities.Product, int64, error) {
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

	if price > 0 {
		switch price {
		case 1:
			query = query.Where("price < ?", 1000)
		case 2:
			query = query.Where("price >= ? AND price < ?", 1000, 5000)
		case 3:
			query = query.Where("price >= ?", 5000)
		}
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
