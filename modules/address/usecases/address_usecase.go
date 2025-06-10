package usecases

import (
	"e-com/modules/address/repositories"
	"e-com/modules/address/reqaddress"
	"e-com/modules/entities"
	"errors"
)

type AddressUsecase interface {
	Create(userID uint, data reqaddress.Reqaddress) (*entities.Address, error)
	FindAll(userId uint) ([]entities.Address, error)
}

type addressUsecase struct {
	repo repositories.AddressRepository
}

func NewAddressUsecase(repo repositories.AddressRepository) AddressUsecase {
	return &addressUsecase{repo: repo}
}

func (u *addressUsecase) Create(userID uint, data reqaddress.Reqaddress) (*entities.Address, error) {
	if data.RecipientName == "" || data.Province == "" {
		return nil, errors.New("recipient name and province are required")
	}

	if data.IsDefault {
		user, err := u.repo.FindDefaultAddressByUserID(userID)
		if err != nil {
			return nil, err
		}

		if user != nil {
			if err := u.repo.UnsetDefaultAddress(userID); err != nil {
				return nil, err
			}
		}
	}

	address := &entities.Address{
		UserID:        userID,
		Type:          data.Type,
		RecipientName: data.RecipientName,
		LastName:      data.LastName,
		Phone:         data.Phone,
		Province:      data.Province,
		District:      data.District,
		SubDistrict:   data.SubDistrict,
		PostalCode:    data.PostalCode,
		IsDefault:     data.IsDefault,
	}

	if err := u.repo.Create(address); err != nil {
		return nil, err
	}
	return address, nil
}

func (u *addressUsecase) FindAll(userId uint) ([]entities.Address, error) {
	data, err := u.repo.FindAll(userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
