package product

import (
	"e-com/modules/product/controllers"
	"e-com/modules/product/repositories"
	"e-com/modules/product/usecases"

	"gorm.io/gorm"
)

type ProductModule struct {
	Controller *controllers.ProductController
}

func New(db *gorm.DB) *ProductModule {
	productRepo := repositories.NewProductRepository(db)
	productUsecase := usecases.NewProductUsecase(productRepo)
	productController := controllers.NewProductController(productUsecase)

	return &ProductModule{
		Controller: productController,
	}
}
