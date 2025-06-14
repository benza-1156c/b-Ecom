package controllers

import (
	"e-com/modules/cart/reqcart"
	"e-com/modules/cart/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CartController struct {
	usecases usecases.CartUsecase
}

func NewCartController(usecases usecases.CartUsecase) *CartController {
	return &CartController{usecases: usecases}
}

func (cc *CartController) Create(c *fiber.Ctx) error {
	userid := c.Locals("userID").(uint)
	var req reqcart.ReqCart
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	err := cc.usecases.Create(userid, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func (cc *CartController) FindAllCartItemByCartid(c *fiber.Ctx) error {
	userid := c.Locals("userID").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	data, err := cc.usecases.FindAllCartItemByCartid(userid, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})

}
