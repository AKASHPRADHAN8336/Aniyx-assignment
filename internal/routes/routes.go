package routes

import (
	"github.com/AKASHPRADHAN8336/aniyxProject/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	userGroup := app.Group("/users")
	{
		userGroup.Post("/", h.CreateUser)
		userGroup.Get("/:id", h.GetUser)
		userGroup.Get("/", h.ListUsers)
		userGroup.Put("/:id", h.UpdateUser)
		userGroup.Delete("/:id", h.DeleteUser)
	}
}
