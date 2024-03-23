package middleware

import (
	"crypto/rsa"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	jwt2 "github.com/kmesiab/equilibria-api/lib/jwt"
	"github.com/kmesiab/equilibria-api/models"
)

var tokenService = jwt2.TokenService{}
var rotator = jwt2.KeyRotator{}

// Authorize This function is responsible for validating the incoming requests
// and ensuring that the user is authorized to access the resources.
func Authorize(c *fiber.Ctx) error {

	var err error
	var tokenString string
	var publicKey *rsa.PublicKey
	var customClaims *jwt2.CustomClaims

	// Get the public key
	if publicKey, err = tokenService.GetPublicKey(&rotator); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get the public key",
		})
	}

	// Get the bearer token
	if tokenString, err = getBearerToken(c); err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or invalid bearer token",
		})
	}

	// Test if signed correctly
	if customClaims, err = tokenService.Validate(tokenString, publicKey); err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": fmt.Sprintf("Invalid bearer token, %s", err.Error()),
		})
	}

	// Test if expired
	if IsTokenExpired(customClaims.ExpiresAt) {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": fmt.Sprintf("Expired bearer token, %s", err.Error()),
		})
	}

	// Test if an account is active.
	if customClaims.AccountStatus != "Active" {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Account not active",
		})
	}

	status := models.StringToAccountStatus(customClaims.AccountStatus)

	user := models.User{
		ID:              customClaims.UserID,
		Password:        nil,
		PhoneNumber:     customClaims.PhoneNumber,
		PhoneVerified:   customClaims.PhoneVerified,
		Firstname:       customClaims.Firstname,
		Lastname:        customClaims.Lastname,
		Email:           customClaims.Email,
		AccountStatusID: customClaims.AccountStatusID,
		AccountStatus:   status,
	}

	c.Locals("user", user)

	// Call next middleware
	if err = c.Next(); err != nil {

		return err
	}

	return nil
}

func IsTokenExpired(expiresAt int64) bool {
	now := time.Now().Unix()
	return now > expiresAt
}

func getBearerToken(c *fiber.Ctx) (string, error) {
	bearer := c.Get("Authorization")

	if bearer == "" {

		return "", fmt.Errorf("bearer token not found")
	}

	tokenString := strings.TrimPrefix(bearer, "Bearer ")

	return tokenString, nil
}
