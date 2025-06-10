package middlewares

import (
	"e-com/modules/entities"
	"e-com/pkg/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthCheck(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing token",
		})
	}

	claims, err := utils.ParsedToken(token)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid user exp",
		})
	}

	if int64(exp) < (time.Now().Unix()) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token expired",
		})
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	email, ok := claims["email"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid user Email",
		})
	}

	c.Locals("userID", uint(userID))
	c.Locals("email", email)

	return c.Next()
}

func Refresh_Token(c *fiber.Ctx) error {
	token := c.Cookies("refresh_token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing token",
		})
	}

	claims, err := utils.ParsedToken(token)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid user exp",
		})
	}

	if int64(exp) < (time.Now().Unix()) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token expired",
		})
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	c.Locals("userID", uint(userID))

	return c.Next()
}

func AdminCheck(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("access_token")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing token",
			})
		}

		claims, err := utils.ParsedToken(token)
		if err != nil {
			log.Println(err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user exp",
			})
		}

		if int64(exp) < (time.Now().Unix()) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "token expired",
			})
		}

		userID, ok := claims["id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user id",
			})
		}

		email, ok := claims["email"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user Email",
			})
		}

		c.Locals("userID", uint(userID))
		c.Locals("email", email)

		var admin entities.User
		if err := db.Where("id = ?", userID).First(&admin).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		if admin.Role != "admin" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "noๆ",
			})
		}

		return c.Next()
	}
}
