package routes

import (
	"backend-project/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the User Management API",
			"endpoints": []string{"/users (GET, POST)", "/users/:id (GET, PUT, DELETE)"},
		})
	})

	api := app.Group("/users")

	api.Post("/", userHandler.CreateUser)
	api.Get("/:id", userHandler.GetUserByID)
	api.Put("/:id", userHandler.UpdateUser)
	api.Delete("/:id", userHandler.DeleteUser)
	api.Get("/", userHandler.ListUsers)
}
