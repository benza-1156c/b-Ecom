package cart

import (
	"e-com/modules/cart/controllers"
	"e-com/modules/cart/repositories"
	"e-com/modules/cart/usecases"
	"e-com/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CartModule struct {
	Controller *controllers.CartController
}

func New(app *fiber.App, db *gorm.DB) {
	cartRepo := repositories.NewCartRepository(db)
	cartUsecase := usecases.NewCartUsecase(cartRepo)
	cartController := controllers.NewCartController(cartUsecase)

	cartGroup := app.Group("/api/cart")

	cartGroup.Post("/", middlewares.AuthCheck, cartController.Create)
	cartGroup.Get("/:id", middlewares.AuthCheck, cartController.FindAllCartItemByCartid)
}
