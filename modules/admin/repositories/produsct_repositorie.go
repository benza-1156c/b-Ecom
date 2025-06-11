package repositories

import (
	"e-com/modules/entities"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(data *entities.Product) error
	CreateImages(productID uint, urls []string) error
	FindByReq(filter, search string, page, limit, category, brand, price int) ([]entities.Product, error)
	FindOneById(id uint) (*entities.Product, error)
	Update(id uint, data *entities.Product) error
	DeleteProduct(id uint) error
	FindTotal() (int64, error)
	// Image
	FindOneImageByID(id uint) (*entities.Image, error)
	DeleteImage(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(data *entities.Product) error {
	return r.db.Create(data).Error
}

func (r *productRepository) CreateImages(productID uint, urls []string) error {
	var images []entities.Image
	for _, url := range urls {
		images = append(images, entities.Image{
			Url:       url,
			ProductID: productID,
		})
	}
	return r.db.Create(&images).Error
}

func (r *productRepository) FindByReq(filter, search string, page, limit, categoryID, brandID, price int) ([]entities.Product, error) {
	var products []entities.Product
	query := r.db.Preload("Images").Preload("Categories").Preload("Brand")

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	if brandID > 0 {
		query = query.Where("brand_id = ?", brandID)
	}

	if price > 0 {
		switch price {
		case 1:
			query = query.Where("price < ?", 10000)
		case 2:
			query = query.Where("price >= ? AND price < ?", 10000, 25000)
		case 3:
			query = query.Where("price >= ? AND price < ?", 25000, 50000)
		case 4:
			query = query.Where("price >= ?", 50000)
		}
	}
	switch filter {
	case "name":
		query = query.Order("name ASC")
	case "newest":
		query = query.Order("created_at DESC")
	case "popular":
		query = query.Order("sold DESC")
	default:
		query = query.Order("created_at DESC")
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindOneById(id uint) (*entities.Product, error) {
	product := &entities.Product{}
	if err := r.db.Preload("Images").Where("id", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Update(id uint, data *entities.Product) error {
	var product entities.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return err
	}

	product.Name = data.Name
	product.Description = data.Description
	product.Price = data.Price
	product.Count = data.Count
	product.Featured = data.Featured
	product.Status = data.Status
	product.Sku = data.Sku
	product.CategoriesID = data.CategoriesID
	product.BrandID = data.BrandID

	return r.db.Save(&product).Error
}

func (r *productRepository) FindOneImageByID(id uint) (*entities.Image, error) {
	product := &entities.Image{}
	if err := r.db.Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) FindAllImageByproductID(id uint) ([]entities.Image, error) {
	var product []entities.Image
	if err := r.db.Where("product_id", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) DeleteImage(id uint) error {
	product := &entities.Image{}
	if err := r.db.Delete(product, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *productRepository) DeleteProduct(id uint) error {
	product := &entities.Product{}
	if err := r.db.Delete(product, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *productRepository) FindTotal() (int64, error) {
	var total int64
	if err := r.db.Model(&entities.Product{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
