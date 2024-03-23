package middleware

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This test assumes you have a method to generate a valid JWT for testing purposes.
func TestAuthorizeWithValidToken(t *testing.T) {
	// Setup Fiber app and middleware
	app := fiber.New()
	app.Use(Authorize)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Generate a token for testing. This should be replaced with your actual token generation logic.
	// For simplicity, this example assumes the token is directly accessible and valid.
	testToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJlcXVpbGlicmlhIiwiZXhwIjoxNzExMTgxMjMyLCJpc3MiOiJlcXVpbGlicmlhIiwiVXNlcklEIjozNiwiRW1haWwiOiJrbWVzaWFiK2VxdWlsaWJyaWFfc21zQGdtYWlsLmNvbSIsIlBob25lTnVtYmVyIjoiKzEyNTMzMjQzMDcxIiwiRmlyc3RuYW1lIjoiS2V2aW4iLCJMYXN0bmFtZSI6Ik1lc2lhYiIsIkFjY291bnRTdGF0dXMiOiJBY3RpdmUiLCJBY2NvdW50U3RhdHVzSUQiOjIsIlBob25lVmVyaWZpZWQiOnRydWV9.r6Asn3ZsFbLSG2AbK70w2BVxHxzHqiiETlYR7Hgb0hHEzLln6TXFiZOtPtKu36qIXPtfGAhuU4uGLBwKrQ1hNbPwDSG26lZhyM0Ox94WuDbcNGdGfPcKwFZAUhJ8_5XWVbNnHcDlNPif3e5fDMImdvwwoukiceK5WTv9ToZNMB7WPSIcExAN_AvKqd1ux6nMb59mvHmNRYZWYCKXP4nFcdZwgkXuNBPlCpMY4rvBBg4ks3k-qhvPEtEne_yVAPCwPg7jTfmYl0rYiyZF_8Rg6EpBSpA_Yjc6SG9XvZd-KETDNX8YuJ36qDMYU2rDUGt2JK9DwmAHyZgSsJ_Vi02zcAsgLu0vkObB5fW4mrhoIOU8EiE_1HSj1ZSQpku11K26IqYg15MPRIfZdDJdjLL5o74kMev-N2xvmyehFCsTj4t1KKLKU74YAPWBfOqSHE0jcDbbtRAxYKpE--SbI6_2ve_p1XJNoybIkocdFTyUhILdgEdBXnL-vMRxx05zqd8QIR16ewwBDqKh9lEUaolAiBVwdF4YpqcRXORabEdpVec9MK9GBhHrbpglb31rviPE5u7Pbs5WSoLmY9IZLgc5J7VKTievOR2mZ1HLxkOv7NSIy1jBTrBDhLqf4KIC_mGZAuo2c5QxPsg5xs1u-wl1d1f04yvhYo5qxJs-SpRSMB0"

	// Create a request with the Authorization header
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)

	// Test the request
	resp, err := app.Test(req, -1)
	assert.Error(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

// This test assumes you have a method to generate a valid JWT for testing purposes.
func TestAuthorizeWithInValidToken(t *testing.T) {
	// Setup Fiber app and middleware
	app := fiber.New()
	app.Use(Authorize)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Generate a token for testing. This should be replaced with your actual token generation logic.
	// For simplicity, this example assumes the token is directly accessible and valid.
	testToken := "invalid-token"

	// Create a request with the Authorization header
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)

	// Test the request
	resp, err := app.Test(req, -1)

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

// This test assumes you have a method to generate a valid JWT for testing purposes.
func TestAuthorizeWithNoToken(t *testing.T) {
	// Setup Fiber app and middleware
	app := fiber.New()
	app.Use(Authorize)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Create a request with the Authorization header
	req := httptest.NewRequest("GET", "/", nil)

	// Test the request
	resp, err := app.Test(req, -1)

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

func TestIsTokenExpired(t *testing.T) {
	// Arrange
	now := time.Now().Unix()
	expiresAt := now - 1000

	// Act
	result := IsTokenExpired(expiresAt)

	// Assert
	assert.Equal(t, true, result)
}

func TestIsTokenNotExpired(t *testing.T) {
	// Arrange
	now := time.Now().Unix()
	expiresAt := now + 10000

	// Act
	result := IsTokenExpired(expiresAt)

	// Assert
	assert.Equal(t, false, result)
}
