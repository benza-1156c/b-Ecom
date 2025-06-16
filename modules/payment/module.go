package payment

import (
	"e-com/modules/payment/controllers"
	"e-com/modules/payment/usecases"

	"github.com/gofiber/fiber/v2"
)

type PaymentModule struct {
	Controller *controllers.PayMentController
}

func NewModulePayment(app *fiber.App) {
	paymentUsecase := usecases.NewPaymentUsacase()
	paymentController := controllers.NewPayMentController(paymentUsecase)

	paymentGroup := app.Group("/api/payment")

	paymentGroup.Post("/promptpay", paymentController.GenerateQR)

}
