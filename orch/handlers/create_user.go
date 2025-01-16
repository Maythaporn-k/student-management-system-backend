package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// TODO:Create a new student
func CreateUser(c *fiber.Ctx) error {
	var student InsertStudent

	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	studentJSON, err := json.Marshal(student)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to serialize student data")
	}

	resp, err := sendRequestToCoreValidate("POST", "/core/create-user", bytes.NewReader(studentJSON))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error Internal Server")
	}
	defer resp.Body.Close()

	body, err := readResponseBodyValidate(resp)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Failed to read response body")
	}
	return c.Status(resp.StatusCode).Send(body)
}
