package controllers

import (
	"e-com/modules/admin/req"
	"e-com/modules/admin/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecases usecases.UserUsecase
}

func NewUserController(usecases usecases.UserUsecase) *UserController {
	return &UserController{usecases: usecases}
}

func (cc *UserController) GetAll(c *fiber.Ctx) error {
	search := c.Query("search")
	role := c.Query("role")

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	users, total, err := cc.usecases.GetAll(role, search, page, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    users,
		"total":   total,
	})
}

func (cc *UserController) Delete(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	if err := cc.usecases.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func (cc *UserController) UpdateStatus(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	var req req.ReqUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := cc.usecases.UpdateStatus(uint(id), req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func (cc *UserController) FindTotal(c *fiber.Ctx) error {
	data, err := cc.usecases.FindTotal()
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
