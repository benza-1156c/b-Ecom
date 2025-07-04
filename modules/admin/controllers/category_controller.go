package controllers

import (
	"e-com/modules/admin/req"
	"e-com/modules/admin/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	usecases usecases.CategoryUsecase
}

func NewCategoryController(usecases usecases.CategoryUsecase) *CategoryController {
	return &CategoryController{usecases: usecases}
}

func (cc *CategoryController) Create(c *fiber.Ctx) error {
	var req req.ReqCategory

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	Category, err := cc.usecases.Create(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    Category,
	})
}

func (cc *CategoryController) Update(c *fiber.Ctx) error {
	var req req.ReqCategory
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	Category, err := cc.usecases.Update(uint(id), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    Category,
	})
}

func (cc *CategoryController) FindAll(c *fiber.Ctx) error {
	category, err := cc.usecases.FindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    category,
	})
}

func (cc *CategoryController) Delete(c *fiber.Ctx) error {
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
