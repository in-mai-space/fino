package entities

import (
	"fino/internal/models"
	"fino/internal/utilities"

	"github.com/gofiber/fiber/v2"
)

type UserControllerInterface interface {
	CreateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type UserController struct {
	service UserServiceInterface
}

func NewUserController(service UserServiceInterface) *UserController {
	return &UserController{service: service}
}

func (controller *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	createdUser, err := controller.service.CreateUser(&user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (controller *UserController) GetUser(c *fiber.Ctx) error {
	id, err := utilities.ValidateUUID(c.Locals("userID").(string))

	if err != nil {
		return utilities.NewBadRequestError("Invalid UUID", err.Error())
	}

	user, err := controller.service.GetUser(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (controller *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := utilities.ValidateUUID(c.Locals("userID").(string))
	if err != nil {
		return utilities.NewBadRequestError("Invalid UUID", err.Error())
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utilities.NewBadRequestError("Invalid UUID", err.Error())
	}

	updatedUser, err := controller.service.UpdateUser(id, &user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

func (controller *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := utilities.ValidateUUID(c.Locals("userID").(string))
	if err != nil {
		return utilities.NewBadRequestError("Invalid UUID", err.Error())
	}

	if err := controller.service.DeleteUser(id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
