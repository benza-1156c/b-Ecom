package usecases

import (
	"e-com/modules/entities"
	"e-com/modules/product/repositories"
)

type ProductUsecase interface {
	FindAll(search string, page, limit, categoryID, brandID, minPrice, maxPrice int) ([]entities.Product, int64, error)
}

type productUsecase struct {
	repo repositories.ProductRepository
}

func NewProductUsecase(repo repositories.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (u *productUsecase) FindAll(search string, page, limit, categoryID, brandID, minPrice, maxPrice int) ([]entities.Product, int64, error) {
	data, total, err := u.repo.FindAll(search, page, limit, categoryID, brandID, minPrice, maxPrice)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}
