package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// TODO:Get the list of students
func StudentList(c *fiber.Ctx) error {
	resp, err := http.Get(fmt.Sprintf("%s/core/student-list", coreBaseURL))
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
