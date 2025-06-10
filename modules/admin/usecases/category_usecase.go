package usecases

import (
	"e-com/modules/admin/repositories"
	"e-com/modules/admin/req"
	"e-com/modules/entities"
)

type CategoryUsecase interface {
	Create(data *req.ReqCategory) (*entities.Category, error)
	Update(id uint, data *req.ReqCategory) (*entities.Category, error)
	FindAll() ([]entities.Category, error)
	Delete(id uint) error
}

type categoryUsecase struct {
	repo repositories.CategoryRepository
}

func NewCategoryUsecase(repo repositories.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo: repo}
}

func (u *categoryUsecase) Create(data *req.ReqCategory) (*entities.Category, error) {
	category := &entities.Category{
		Name: data.Name,
		Icon: data.Icon,
	}

	if err := u.repo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (u *categoryUsecase) Update(id uint, data *req.ReqCategory) (*entities.Category, error) {
	category := &entities.Category{
		Name: data.Name,
		Icon: data.Icon,
	}

	if err := u.repo.Update(id, category); err != nil {
		return nil, err
	}

	category.ID = id
	return category, nil
}

func (u *categoryUsecase) FindAll() ([]entities.Category, error) {
	category, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (u *categoryUsecase) Delete(id uint) error {
	if err := u.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
