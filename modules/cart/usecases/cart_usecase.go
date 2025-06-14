package usecases

import (
	"e-com/modules/cart/repositories"
	"e-com/modules/cart/reqcart"
	"e-com/modules/entities"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CartUsecase interface {
	Create(userid uint, data reqcart.ReqCart) error
	FindAllCartItemByCartid(userid, id uint) ([]entities.CartItem, error)
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
		if userid != user_cart.UserID {
			return errors.New("???")
		}
		newItem := &entities.CartItem{
			CartID:    user_cart.ID,
			ProductID: data.ProductID,
			Quantity:  data.Quantity,
		}
		if err := u.repo.CreateCartItem(newItem); err != nil {
			return err
		}
	} else {
		if userid != user_cart.UserID {
			return errors.New("???")
		}
		cartitem.Quantity += data.Quantity
		if err := u.repo.UpdateCartItem(cartitem); err != nil {
			return err
		}
	}

	return nil
}

func (u *cartUsecase) FindAllCartItemByCartid(userid, id uint) ([]entities.CartItem, error) {
	user, err := u.repo.FindCartByUserId(userid)
	if err != nil {
		return nil, err
	}
	if user.UserID != userid {
		return nil, err
	}

	data, err := u.repo.FindAllCartItemByCartid(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
