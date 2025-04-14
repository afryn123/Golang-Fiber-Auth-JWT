package middleware

import (
    "fiber-auth-app/utils"
    "github.com/gofiber/fiber/v2"
    "strings"
)

func JWTProtected() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
        }

        userID, err := utils.ParseJWT(parts[1])
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
        }

        c.Locals("userID", userID)
        return c.Next()
    }
}
