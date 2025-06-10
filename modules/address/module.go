package address

import (
	"e-com/modules/address/controllers"
	"e-com/modules/address/repositories"
	"e-com/modules/address/usecases"

	"gorm.io/gorm"
)

type AddressModule struct {
	Controllers *controllers.AddressController
}

func NewAddressModule(db *gorm.DB) *AddressModule {
	addressrepo := repositories.NewAddressRepository(db)
	addressusecase := usecases.NewAddressUsecase(addressrepo)
	addressController := controllers.NewAddressController(addressusecase)

	return &AddressModule{
		Controllers: addressController,
	}
}
