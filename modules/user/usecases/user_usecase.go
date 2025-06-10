package usecases

import (
	"e-com/modules/entities"
	"e-com/modules/user/repositories"
	"e-com/pkg/utils"
)

type UserUsecase interface {
	FindOneById(id uint) (*entities.User, error)
	Refresh_Token(id uint) (*entities.User, string, string, error)
}

type userUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) FindOneById(id uint) (*entities.User, error) {
	user, err := u.repo.FindOneById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Refresh_Token(id uint) (*entities.User, string, string, error) {
	user, err := u.repo.FindOneById(id)
	if err != nil {
		return nil, "", "", err
	}

	token, err := utils.CreateTokenJWT(user.ID, &user.Email, 4)
	if err != nil {
		return nil, "", "", err
	}

	refreshToken, err := utils.CreateTokenJWT(user.ID, nil, 150)
	if err != nil {
		return nil, "", "", err
	}

	return user, token, refreshToken, nil
}
