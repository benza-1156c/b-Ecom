package controllers

import (
	"e-com/modules/user/usecases"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecases usecases.UserUsecase
}

func NewUserController(usecases usecases.UserUsecase) *UserController {
	return &UserController{usecases: usecases}
}

func (cc *UserController) Me(c *fiber.Ctx) error {
	id := c.Locals("userID").(uint)
	data, err := cc.usecases.FindOneById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func (cc *UserController) Refresh_Token(c *fiber.Ctx) error {
	id := c.Locals("userID").(uint)
	data, accessToken, refreshToken, err := cc.usecases.Refresh_Token(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   accessToken,
		Expires: time.Now().Add(4 * time.Hour),
		Secure:  false,
	})

	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   refreshToken,
		Expires: time.Now().Add(150 * time.Hour),
		Secure:  false,
	})

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
