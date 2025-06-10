package controllers

import (
	"context"
	"e-com/config"
	"e-com/modules/auth/req"
	"e-com/modules/auth/usecases"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	oauth2api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type AuthController struct {
	usecases usecases.AuthUsecase
}

func NewAuthController(u usecases.AuthUsecase) *AuthController {
	return &AuthController{usecases: u}
}

func (cc *AuthController) HandleGoogleLogin(c *fiber.Ctx) error {
	url := config.GoogleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func (cc *AuthController) HandleGoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Code not found")
	}

	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Token exchange failed")
	}

	client := config.GoogleOAuthConfig.Client(context.Background(), token)

	ctx := context.Background()
	oauth2srv, err := oauth2api.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create OAuth2 service")
	}

	userinfo, err := oauth2api.NewUserinfoService(oauth2srv).V2.Me.Get().Do()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info")
	}

	data := &req.ReqAuth{
		Email:    userinfo.Email,
		UserName: userinfo.Name,
		Avater:   userinfo.Picture,
	}

	_, accessToken, refreshToken, err := cc.usecases.CreateUser(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   accessToken,
		Expires: time.Now().Add(4 * time.Hour),
		Secure:  false,
	})

	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   refreshToken,
		Expires: time.Now().Add(150 * time.Hour),
		Secure:  false,
	})

	return c.Redirect("http://localhost:3000")
}
