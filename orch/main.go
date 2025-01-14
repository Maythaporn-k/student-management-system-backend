package main

import (
	"orch/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/orch/student-list", handlers.StudentList)
	app.Post("/orch/create-user", handlers.CreateStudent)
	app.Delete("/orch/delete-user", handlers.DeleteStudent)
	app.Put("/orch/edit-user", handlers.EditStudent)

	app.Listen(":3001")
}
