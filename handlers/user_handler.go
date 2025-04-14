package handlers

import (
	"fiber-auth-app/repository"
	"fiber-auth-app/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Profile(c *fiber.Ctx) error
}

type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) UserHandler {
	return &userHandler{
		repo: repo,
	}
}

func (h userHandler) Profile(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return utils.JSONError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	idInt, ok := userID.(uint)
	if !ok {
		return utils.JSONError(c, fiber.StatusBadGateway, "Invalid user ID format")
	}

	uid := int(idInt)

	user, err := h.repo.FindById(uid)
	if err != nil {
		return utils.JSONError(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	data := fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}
	return utils.JSONSuccess(c, data, "Success")
}
