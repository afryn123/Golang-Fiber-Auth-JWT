package handlers

import (
	"fiber-auth-app/models"
	"fiber-auth-app/repository"
	"fiber-auth-app/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	repo     repository.UserRepository
	validate *validator.Validate
}

// Constructor
func NewAuthHandler(repo repository.UserRepository) AuthHandler {
	return &authHandler{
		repo:     repo,
		validate: validator.New(),
	}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return utils.JSONError(c, fiber.StatusBadRequest, err.Error())
	}

	hashedPassword, _ := utils.HashPassword(input.Password)
	input.Password = hashedPassword

	if err := h.repo.Create(&input); err != nil {
		return utils.JSONError(c, fiber.StatusInternalServerError, "Could not create user")
	}

	return utils.JSONSuccess(c, nil, "Success to create data")
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return utils.JSONError(c, fiber.StatusUnauthorized, err.Error())
	}

	user, err := h.repo.FindByEmail(input.Email)
	if err != nil {
		return utils.JSONError(c, fiber.StatusUnauthorized, "Invalid Email")
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return utils.JSONError(c, fiber.StatusUnauthorized, "Invalid Password")
	}

	token, _ := utils.GenerateJWT(user.ID)
	return utils.JSONSuccess(c, fiber.Map{"token": token}, "Success")
}
