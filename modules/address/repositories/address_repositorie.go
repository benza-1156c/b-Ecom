package repositories

import (
	"e-com/modules/entities"
	"errors"

	"gorm.io/gorm"
)

type AddressRepository interface {
	FindDefaultAddressByUserID(userid uint) (*entities.Address, error)
	FindAll(userId uint) ([]entities.Address, error)
	FindOneByid(id uint) (*entities.Address, error)
	Create(data *entities.Address) error
	Update(id uint, data *entities.Address) error
	UnsetDefaultAddress(userID uint) error
	Delete(userID, id uint) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) FindAll(userId uint) ([]entities.Address, error) {
	var address []entities.Address
	if err := r.db.Where("user_id = ?", userId).Find(&address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *addressRepository) FindDefaultAddressByUserID(userid uint) (*entities.Address, error) {
	var address entities.Address
	if err := r.db.Where("user_id = ? AND is_default = ?", userid, true).First(&address).Error; err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *addressRepository) FindOneByid(id uint) (*entities.Address, error) {
	address := &entities.Address{}
	if err := r.db.First(address, id).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *addressRepository) Create(data *entities.Address) error {
	return r.db.Create(data).Error
}

func (r *addressRepository) Update(id uint, data *entities.Address) error {
	return r.db.Model(&entities.Address{}).Where("id = ?", id).Updates(data).Error
}

func (r *addressRepository) UnsetDefaultAddress(userID uint) error {
	result := r.db.Model(&entities.Address{}).Where("user_id = ? AND is_default = ?", userID, true).Update("is_default", false)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *addressRepository) Delete(userID, id uint) error {
	var addr entities.Address
	if err := r.db.First(&addr, id).Error; err != nil {
		return err
	}

	if addr.UserID != userID {
		return errors.New("?????")
	}

	return r.db.Delete(&addr).Error
}
