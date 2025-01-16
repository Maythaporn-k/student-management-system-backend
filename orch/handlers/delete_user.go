package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// TODO: Delete a student
func DeleteUser(c *fiber.Ctx) error {
	var studentId DeleteStudent

	if err := c.BodyParser(&studentId); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	studentJSON, err := json.Marshal(map[string]int{"id": studentId.ID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to serialize student data")
	}

	resp, err := sendRequestToCoreValidate("DELETE", "/core/delete-user", bytes.NewReader(studentJSON))
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
