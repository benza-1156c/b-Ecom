package usecases

import (
	"e-com/modules/admin/repositories"
	"e-com/modules/admin/req"
	"e-com/modules/entities"
)

type BrandUsecase interface {
	Create(data *req.ReqBrand) (*entities.Brand, error)
	FindAll() ([]entities.Brand, error)
	Delete(id uint) error
	Update(id uint, data *req.ReqBrand) (*entities.Brand, error)
}

type brandUsecase struct {
	repo repositories.BrandRepository
}

func NewBrandUsecase(repo repositories.BrandRepository) BrandUsecase {
	return &brandUsecase{repo: repo}
}

func (u *brandUsecase) Create(data *req.ReqBrand) (*entities.Brand, error) {
	brand := &entities.Brand{
		Name: data.Name,
		Icon: data.Icon,
	}

	err := u.repo.Create(brand)
	if err != nil {
		return nil, err
	}

	return brand, nil
}

func (u *brandUsecase) FindAll() ([]entities.Brand, error) {
	data, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *brandUsecase) Delete(id uint) error {
	if err := u.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *brandUsecase) Update(id uint, data *req.ReqBrand) (*entities.Brand, error) {
	brand := &entities.Brand{
		Name: data.Name,
		Icon: data.Icon,
	}

	if err := u.repo.Update(id, brand); err != nil {
		return nil, err
	}
	brand.ID = id
	return brand, nil
}
