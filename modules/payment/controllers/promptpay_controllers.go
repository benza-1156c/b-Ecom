package controllers

import (
	"e-com/modules/payment/usecases"

	"github.com/gofiber/fiber/v2"
)

type PayMentController struct {
	usecases usecases.PaymentUsacase
}

func NewPayMentController(usecases usecases.PaymentUsacase) *PayMentController {
	return &PayMentController{usecases: usecases}
}

type PromptPayRequest struct {
	Amount string `json:"amount"`
}

func (cc *PayMentController) GenerateQR(c *fiber.Ctx) error {
	var req PromptPayRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	payload := cc.usecases.GeneratePromptPayPayload("0650269486", req.Amount)

	return c.JSON(fiber.Map{
		"success": true,
		"payload": payload,
	})
}
