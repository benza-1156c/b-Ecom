package controllers

import (
	"e-com/modules/address/reqaddress"
	"e-com/modules/address/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AddressController struct {
	usecases usecases.AddressUsecase
}

func NewAddressController(usecases usecases.AddressUsecase) *AddressController {
	return &AddressController{usecases: usecases}
}

func (cc *AddressController) Create(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	var req reqaddress.Reqaddress
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	data, err := cc.usecases.Create(userId, req)
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

func (cc *AddressController) Update(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	userId := c.Locals("userID").(uint)
	var req reqaddress.Reqaddress
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	data, err := cc.usecases.Update(userId, uint(id), req)
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

func (cc *AddressController) FindAll(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	data, err := cc.usecases.FindAll(userId)
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

func (cc *AddressController) Delete(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	userId := c.Locals("userID").(uint)

	if err := cc.usecases.Delete(userId, uint(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
