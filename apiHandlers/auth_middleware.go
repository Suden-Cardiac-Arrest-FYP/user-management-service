package apiHandlers

import (
	"fmt"
	"User-Mgt/dto"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/url"
	"strings"
	"time"
)

type AuthMiddleware struct {
	config dto.AuthConfig
}

func NewAuthMiddleware(config dto.AuthConfig) *AuthMiddleware {
	return &AuthMiddleware{
		config: config,
	}
}

func (a *AuthMiddleware) ValidateToken(c *fiber.Ctx) error {
	issuerURL, err := url.Parse("https://" + a.config.AUTH0_DOMAIN + "/")
	if err != nil {
		log.Fatalf("failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{a.config.AUTH0_AUDIENCE},
	)
	if err != nil {
		log.Fatalln("failed to set up the jwt validator")
	}

	authHeader := c.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid authorization header",
		})
	}

	_, err = jwtValidator.ValidateToken(c.Context(), authHeaderParts[1])
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	return c.Next()
}
