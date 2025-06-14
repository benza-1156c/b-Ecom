package controllers

import (
	"e-com/modules/product/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	usecases usecases.ProductUsecase
}

func NewProductController(usecases usecases.ProductUsecase) *ProductController {
	return &ProductController{usecases: usecases}
}

func (cc *ProductController) FindAll(c *fiber.Ctx) error {
	search := c.Query("search")
	category, _ := strconv.Atoi(c.Query("category", "0"))
	brand, _ := strconv.Atoi(c.Query("brand", "0"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "12"))
	minPrice, _ := strconv.Atoi(c.Query("minPrice", "0"))
	maxPrice, _ := strconv.Atoi(c.Query("maxPrice", "0"))

	products, total, err := cc.usecases.FindAll(search, page, limit, category, brand, minPrice, maxPrice)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	hasMore := (int64(page) * int64(limit)) < total

	return c.JSON(fiber.Map{
		"success":     true,
		"data":        products,
		"total":       total,
		"hasMore":     hasMore,
		"currentPage": page,
	})
}

func (cc *ProductController) FindOneById(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	data, err := cc.usecases.FindOneById(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})

}

func (cc *ProductController) FindProductFeatured(c *fiber.Ctx) error {
	data, err := cc.usecases.FindProductFeatured()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func (cc *ProductController) FindAllByCategory(c *fiber.Ctx) error {
	categoryidquy := c.Query("category")
	limitqy := c.Query("limit")
	categoryid, err := strconv.Atoi(categoryidquy)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	limit, err := strconv.Atoi(limitqy)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	data, err := cc.usecases.FindAllByCategory(uint(categoryid), uint(limit))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
