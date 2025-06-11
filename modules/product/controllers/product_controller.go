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
