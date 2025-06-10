package usecases

import (
	"e-com/modules/auth/repositories"
	"e-com/modules/auth/req"
	"e-com/modules/entities"
	"e-com/pkg/utils"
	"time"

	"gorm.io/gorm"
)

type AuthUsecase interface {
	CreateUser(data *req.ReqAuth) (*entities.User, string, string, error)
}

type authUsecase struct {
	repo repositories.AuthRepository
}

func NewAuthUsecase(repo repositories.AuthRepository) AuthUsecase {
	return &authUsecase{repo: repo}
}

func (u *authUsecase) CreateUser(data *req.ReqAuth) (*entities.User, string, string, error) {
	exuser, err := u.repo.FindOneUserByemail(data.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, "", "", err
	}

	user := &entities.User{
		UserName:  data.UserName,
		Email:     data.Email,
		Avatar:    data.Avater,
		Role:      "user",
		Status:    "active",
		CreatedAt: time.Now(),
	}

	if exuser != nil {
		user.ID = exuser.ID
		err := u.repo.UpdateUser(user)
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

	err = u.repo.CreateUser(user)
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
