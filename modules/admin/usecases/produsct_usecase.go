package usecases

import (
	"e-com/modules/admin/repositories"
	"e-com/modules/admin/req"
	"e-com/modules/entities"
	"errors"
	"fmt"
)

type ProductUsecase interface {
	Create(data *req.ReqProduct) (*entities.Product, error)
	FindByReq(filter, search string, page, limit, category, brand, price int) ([]entities.Product, error)
	FindOneById(id uint) (*entities.Product, error)
	Update(id uint, data *req.ReqProduct) (*entities.Product, error)
	DeleteProduct(id uint) error
	FindTotal() (int64, error)
	// Image
	FindImage(id uint) (*entities.Image, error)
	DeleteImage(id uint) error
}

type productUsecase struct {
	repo repositories.ProductRepository
}

func NewProductUsecase(repo repositories.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (u *productUsecase) Create(data *req.ReqProduct) (*entities.Product, error) {
	if data.Brand == 0 || data.Category == 0 || data.Count == 0 || data.Description == "" || data.Name == "" || data.Price == 0 || data.Sku == "" {
		return nil, errors.New("ข้อมูลไม่ครบถ้วน")
	}

	product := &entities.Product{
		Name:         data.Name,
		Description:  data.Description,
		Status:       data.Status,
		BrandID:      data.Brand,
		Sku:          data.Sku,
		Count:        data.Count,
		Price:        data.Price,
		Featured:     data.Featured,
		CategoriesID: data.Category,
	}

	if err := u.repo.Create(product); err != nil {
		return nil, err
	}

	if len(data.Images) > 0 {
		err := u.repo.CreateImages(product.ID, data.Images)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (u *productUsecase) FindByReq(filter, search string, page, limit, category, brand, price int) ([]entities.Product, error) {
	data, err := u.repo.FindByReq(filter, search, page, limit, category, brand, price)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *productUsecase) FindOneById(id uint) (*entities.Product, error) {
	data, err := u.repo.FindOneById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *productUsecase) Update(id uint, data *req.ReqProduct) (*entities.Product, error) {
	if data.Brand == 0 || data.Category == 0 || data.Count == 0 || data.Description == "" || data.Name == "" || data.Price == 0 || data.Sku == "" {
		return nil, errors.New("ข้อมูลไม่ครบถ้วน")
	}

	product := &entities.Product{
		Name:         data.Name,
		Description:  data.Description,
		Status:       data.Status,
		BrandID:      data.Brand,
		Sku:          data.Sku,
		Count:        data.Count,
		Price:        data.Price,
		Featured:     data.Featured,
		CategoriesID: data.Category,
	}

	if err := u.repo.Update(id, product); err != nil {
		return nil, err
	}

	fmt.Println(data.Images)
	product.ID = id

	if len(data.Images) > 0 {
		err := u.repo.CreateImages(product.ID, data.Images)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (u *productUsecase) DeleteProduct(id uint) error {
	if err := u.repo.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) FindImage(id uint) (*entities.Image, error) {
	data, err := u.repo.FindOneImageByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *productUsecase) DeleteImage(id uint) error {
	if err := u.repo.DeleteImage(id); err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) FindTotal() (int64, error) {
	data, err := u.repo.FindTotal()
	if err != nil {
		return 0, err
	}
	return data, nil
}
