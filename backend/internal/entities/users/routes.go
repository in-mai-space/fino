package entities

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(api *fiber.Router, db *gorm.DB) {
	userGroup := (*api).Group("/users")

	userTransaction := NewUserTransaction(db)
	userService := NewUserService(userTransaction)
	userController := NewUserController(userService)

	userGroup.Post("/users", userController.CreateUser)
	userGroup.Get("/users/me", userController.GetUser)
	userGroup.Patch("/users/me", userController.UpdateUser)
	userGroup.Delete("/users/me", userController.DeleteUser)
}
