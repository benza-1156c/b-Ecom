package usecases

import (
	"e-com/modules/cart/repositories"
	"e-com/modules/cart/reqcart"
	"e-com/modules/entities"
)

type CartUsecase interface {
	Create(userid uint, data reqcart.ReqCart) error
}

type cartUsecase struct {
	repo repositories.CartRepository
}

func NewCartUsecase(repo repositories.CartRepository) CartUsecase {
	return &cartUsecase{repo: repo}
}

func (u *cartUsecase) Create(userid uint, data reqcart.ReqCart) error {
	cart := &entities.Cart{
		UserID:    userid,
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
	}

	if err := u.repo.Create(cart); err != nil {
		return err
	}
	return nil
}
