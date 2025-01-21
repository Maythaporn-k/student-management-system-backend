package main

import (
	"orch/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// *set content type
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	// middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/orch/student-list", handlers.StudentList)
	app.Post("/orch/create-user", handlers.CreateUser)
	app.Delete("/orch/delete-user", handlers.DeleteUser)
	app.Put("/orch/edit-user", handlers.EditUser)

	app.Listen(":3001")
}
