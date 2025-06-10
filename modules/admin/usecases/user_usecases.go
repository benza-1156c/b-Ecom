package usecases

import (
	"e-com/modules/admin/repositories"
	"e-com/modules/admin/req"
	"e-com/modules/entities"
)

type UserUsecase interface {
	GetAll(role, search string, page, limit int) ([]entities.User, int64, error)
	Delete(id uint) error
	UpdateStatus(id uint, req req.ReqUser) error
	FindTotal() (int64, error)
}

type userUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetAll(role, search string, page, limit int) ([]entities.User, int64, error) {
	data, total, err := u.repo.GetAll(role, search, page, limit)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (u *userUsecase) Delete(id uint) error {
	if err := u.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) UpdateStatus(id uint, req req.ReqUser) error {
	user := &entities.User{
		Status: req.Status,
	}
	if err := u.repo.UpdateStatus(id, user); err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) FindTotal() (int64, error) {
	count, err := u.repo.FindTotal()
	if err != nil {
		return 0, nil
	}
	return count, nil
}
