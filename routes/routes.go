package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"e-com/modules/address"
	authcontroller "e-com/modules/auth/controllers"
	authrepo "e-com/modules/auth/repositories"
	authusecase "e-com/modules/auth/usecases"
	"e-com/modules/product"

	useradminusercontroller "e-com/modules/admin/controllers"
	useradminepo "e-com/modules/admin/repositories"
	useradminsecase "e-com/modules/admin/usecases"

	catadminusercontroller "e-com/modules/admin/controllers"
	catadminepo "e-com/modules/admin/repositories"
	catadminsecase "e-com/modules/admin/usecases"

	brandadminusercontroller "e-com/modules/admin/controllers"
	brandadminepo "e-com/modules/admin/repositories"
	brandadminsecase "e-com/modules/admin/usecases"

	productdminusercontroller "e-com/modules/admin/controllers"
	productadminepo "e-com/modules/admin/repositories"
	productadminsecase "e-com/modules/admin/usecases"

	usercontroller "e-com/modules/user/controllers"
	userepo "e-com/modules/user/repositories"
	usersecase "e-com/modules/user/usecases"
	"e-com/pkg/middlewares"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("api")

	authrepo := authrepo.NewAuthRepository(db)
	authusecase := authusecase.NewAuthUsecase(authrepo)
	authcontroller := authcontroller.NewAuthController(authusecase)

	app.Get("/auth/google/login", authcontroller.HandleGoogleLogin)
	app.Get("/auth/google/callback", authcontroller.HandleGoogleCallback)

	// Admin

	// User
	useradminepo := useradminepo.NewUserRepository(db)
	useradminsecase := useradminsecase.NewUserUsecase(useradminepo)
	useradminusercontroller := useradminusercontroller.NewUserController(useradminsecase)

	api.Get("/admin/users", middlewares.AdminCheck(db), useradminusercontroller.GetAll)
	api.Get("/admin/users/count", middlewares.AdminCheck(db), useradminusercontroller.FindTotal)
	api.Put("/admin/users/status/:id", middlewares.AdminCheck(db), useradminusercontroller.UpdateStatus)
	api.Delete("/admin/users/:id", middlewares.AdminCheck(db), useradminusercontroller.Delete)
	// User

	// category
	catadminepo := catadminepo.NewCategoryRepository(db)
	catadminsecase := catadminsecase.NewCategoryUsecase(catadminepo)
	catadminusercontroller := catadminusercontroller.NewCategoryController(catadminsecase)

	api.Get("/category", catadminusercontroller.FindAll)
	api.Post("/admin/category", middlewares.AdminCheck(db), catadminusercontroller.Create)
	api.Delete("/admin/category/:id", middlewares.AdminCheck(db), catadminusercontroller.Delete)
	api.Put("/admin/category/:id", middlewares.AdminCheck(db), catadminusercontroller.Update)
	// category

	// Brand
	brandadminepo := brandadminepo.NewBrandRepository(db)
	brandadminsecase := brandadminsecase.NewBrandUsecase(brandadminepo)
	brandadminusercontroller := brandadminusercontroller.NewBrandController(brandadminsecase)

	api.Get("/brand", brandadminusercontroller.FindAll)
	api.Post("/admin/brand", middlewares.AdminCheck(db), brandadminusercontroller.Create)
	api.Delete("/admin/brand/:id", middlewares.AdminCheck(db), brandadminusercontroller.Delete)
	api.Delete("/admin/brand/:id", middlewares.AdminCheck(db), brandadminusercontroller.Delete)
	api.Put("/admin/brand/:id", middlewares.AdminCheck(db), brandadminusercontroller.Update)
	// Brand

	// Product
	productadminepo := productadminepo.NewProductRepository(db)
	productadminsecase := productadminsecase.NewProductUsecase(productadminepo)
	productdminusercontroller := productdminusercontroller.NewProductController(productadminsecase)

	api.Get("/admin/product", middlewares.AdminCheck(db), productdminusercontroller.FindByQuery)
	api.Get("/admin/product/:id", middlewares.AdminCheck(db), productdminusercontroller.FindOneById)
	api.Get("/admin/products/count", middlewares.AdminCheck(db), productdminusercontroller.FindTotal)
	api.Post("/admin/product", middlewares.AdminCheck(db), productdminusercontroller.Create)
	api.Put("/admin/product/:id", middlewares.AdminCheck(db), productdminusercontroller.Update)
	api.Delete("/admin/product/:id", middlewares.AdminCheck(db), productdminusercontroller.DeleteProduct)

	// Image
	api.Delete("/admin/product/image/:id", middlewares.AdminCheck(db), productdminusercontroller.DeleteImage)

	// Product

	// Admin

	// User
	userepo := userepo.NewUserRepository(db)
	usersecase := usersecase.NewUserUsecase(userepo)
	usercontroller := usercontroller.NewUserController(usersecase)

	api.Get("@me", middlewares.AuthCheck, usercontroller.Me)
	api.Get("refresh-token", middlewares.Refresh_Token, usercontroller.Refresh_Token)

	// Address
	AddressModule := address.NewAddressModule(db)
	api.Get("/addresses", middlewares.AuthCheck, AddressModule.Controllers.FindAll)
	api.Post("/addresses", middlewares.AuthCheck, AddressModule.Controllers.Create)
	api.Put("/addresses/:id", middlewares.AuthCheck, AddressModule.Controllers.Update)
	api.Delete("/addresses/:id", middlewares.AuthCheck, AddressModule.Controllers.Delete)
	// Address

	// Product
	productModule := product.New(db)
	api.Get("/products", productModule.Controller.FindAll)
	// Product
	// User

}
