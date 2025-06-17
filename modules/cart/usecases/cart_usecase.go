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
	UpdateCount(userid, CartId, productID uint, quantity int) error
	DeleteCartItem(userid, id uint) error
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

func (u *cartUsecase) UpdateCount(userid, CartId, productID uint, quantity int) error {
	cart, err := u.repo.FindCartByUserId(userid)
	if err != nil {
		return err
	}

	if cart.UserID != userid {
		return errors.New("คุณไม่มีสิทธิ์แก้ไขสินค้านี้")
	}

	cartItem, err := u.repo.FindByCartIDAndProductID(CartId, productID)
	if err != nil {
		return errors.New("ไม่พบสินค้านี้ในตะกร้า")
	}

	cartItem.Quantity = quantity
	if err := u.repo.UpdateCountCartItem(cartItem); err != nil {
		return errors.New("error Save to DB")
	}

	return nil
}

func (u *cartUsecase) DeleteCartItem(userid, id uint) error {
	cart, err := u.repo.FindCartItemById(id)
	if err != nil {
		return err
	}

	if cart.Cart.UserID != userid {
		return errors.New("คุณไม่มีสิทธิ์แก้ไขสินค้านี้")
	}

	if err := u.repo.DeleteCartItemByID(id); err != nil {
		return err
	}

	return nil
}
