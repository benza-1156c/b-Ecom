package usecases

import (
	"e-com/modules/address/repositories"
	"e-com/modules/address/reqaddress"
	"e-com/modules/entities"
	"errors"

	"gorm.io/gorm"
)

type AddressUsecase interface {
	Create(userID uint, data reqaddress.Reqaddress) (*entities.Address, error)
	FindAll(userId uint) ([]entities.Address, error)
	Update(userID, id uint, data reqaddress.Reqaddress) (*entities.Address, error)
	Delete(userID, id uint) error
}

type addressUsecase struct {
	repo repositories.AddressRepository
}

func NewAddressUsecase(repo repositories.AddressRepository) AddressUsecase {
	return &addressUsecase{repo: repo}
}

func (u *addressUsecase) Create(userID uint, data reqaddress.Reqaddress) (*entities.Address, error) {
	if data.RecipientName == "" || data.Province == "" || data.Address == "" {
		return nil, errors.New("กรุณากรอกข้อมูลที่อยู่ให้ครบถ้วน")
	}

	if data.IsDefault {
		user, err := u.repo.FindDefaultAddressByUserID(userID)
		if err != nil && err != gorm.ErrRecordNotFound {
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
		Address:       data.Address,
		RecipientName: data.RecipientName,
		LastName:      data.LastName,
		Phone:         data.Phone,
		Province:      data.Province,
		District:      data.District,
		Other:         data.Other,
		SubDistrict:   data.SubDistrict,
		PostalCode:    data.PostalCode,
		ProvinceId:    data.ProvinceId,
		AmphureId:     data.AmphureId,
		TambonId:      data.TambonId,
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

func (u *addressUsecase) Update(userID, id uint, data reqaddress.Reqaddress) (*entities.Address, error) {
	if data.RecipientName == "" || data.Province == "" || data.Address == "" {
		return nil, errors.New("กรุณากรอกข้อมูลที่อยู่ให้ครบถ้วน")
	}

	seaddress, err := u.repo.FindOneByid(id)
	if err != nil {
		return nil, err
	}

	if seaddress.UserID != userID {
		return nil, err
	}

	if data.IsDefault {
		user, err := u.repo.FindDefaultAddressByUserID(userID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		if user != nil {
			if err := u.repo.UnsetDefaultAddress(userID); err != nil {
				return nil, err
			}
		}
	}

	address := &entities.Address{
		Type:          data.Type,
		RecipientName: data.RecipientName,
		LastName:      data.LastName,
		Address:       data.Address,
		Phone:         data.Phone,
		Province:      data.Province,
		District:      data.District,
		Other:         data.Other,
		SubDistrict:   data.SubDistrict,
		PostalCode:    data.PostalCode,
		ProvinceId:    data.ProvinceId,
		AmphureId:     data.AmphureId,
		TambonId:      data.TambonId,
		IsDefault:     data.IsDefault,
	}

	if err := u.repo.Update(id, address); err != nil {
		return nil, err
	}

	address.ID = id
	address.UserID = userID
	return address, nil
}

func (u *addressUsecase) Delete(userID, id uint) error {
	if err := u.repo.Delete(userID, id); err != nil {
		return err
	}
	return nil
}
