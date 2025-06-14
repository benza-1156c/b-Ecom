package usecases

import (
	"e-com/modules/cart/repositories"
	"e-com/modules/cart/reqcart"
	"e-com/modules/entities"
	"time"

	"gorm.io/gorm"
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
	user_cart, err := u.repo.FindCartByUserId(userid)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if user_cart == nil {
		newCart := &entities.Cart{
			UserID:    userid,
			CreatedAt: time.Now(),
		}
		if err := u.repo.Create(newCart); err != nil {
			return err
		}
		user_cart = newCart
	}

	cartitem, _ := u.repo.FindByCartIDAndProductID(userid, data.ProductID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if cartitem == nil {
		newItem := &entities.CartItem{
			CartID:    user_cart.ID,
			ProductID: data.ProductID,
			Quantity:  data.Quantity,
		}
		if err := u.repo.CreateCartItem(newItem); err != nil {
			return err
		}
	} else {
		cartitem.Quantity += data.Quantity
		if err := u.repo.UpdateCartItem(cartitem); err != nil {
			return err
		}
	}

	return nil
}
