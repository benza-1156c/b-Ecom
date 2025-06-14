package usecases

import (
	"e-com/modules/entities"
	"e-com/modules/product/repositories"
)

type ProductUsecase interface {
	FindAll(search string, page, limit, categoryID, brandID, minPrice, maxPrice int) ([]entities.Product, int64, error)
	FindOneById(id uint) (*entities.Product, error)
	FindAllByCategory(id, Limit uint) ([]entities.Product, error)
	FindProductFeatured() ([]entities.Product, error)
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

func (u *productUsecase) FindOneById(id uint) (*entities.Product, error) {
	product, err := u.repo.FindOneById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (u *productUsecase) FindProductFeatured() ([]entities.Product, error) {
	data, err := u.repo.FindProductFeatured()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *productUsecase) FindAllByCategory(id, Limit uint) ([]entities.Product, error) {
	data, err := u.repo.FindAllByCategory(id, Limit)
	if err != nil {
		return nil, err
	}

	return data, nil
}
