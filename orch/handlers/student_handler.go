package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const coreBaseURL = "http://localhost:3002" // Core service URL

func sendRequestToCoreValidate(method, endpoint string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", coreBaseURL, endpoint)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func readResponseBodyValidate(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

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

	switch resp.StatusCode {
	case http.StatusOK:
		return c.Status(fiber.StatusOK).Send(body)
	case http.StatusNotFound:
		return c.Status(fiber.StatusNotFound).SendString("Student list not found")
	default:
		return c.Status(fiber.StatusInternalServerError).SendString("Error Internal Server")
	}

}

// TODO:Create a new student
func CreateStudent(c *fiber.Ctx) error {
	var student struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Grade string `json:"grade"`
	}

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

	switch resp.StatusCode {
	case http.StatusOK:
		return c.Status(fiber.StatusOK).Send(body)
	case http.StatusNotFound:
		return c.Status(fiber.StatusNotFound).SendString("Can't insert new student")
	default:
		return c.Status(fiber.StatusInternalServerError).SendString("Error Internal Server")
	}
}

// TODO: Delete a student
func DeleteStudent(c *fiber.Ctx) error {
	var studentId struct {
		ID int `json:"id"`
	}

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

	switch resp.StatusCode {
	case http.StatusOK:
		return c.Status(fiber.StatusOK).Send(body)
	case http.StatusNotFound:
		return c.Status(fiber.StatusNotFound).SendString("Can't delete student")
	default:
		return c.Status(fiber.StatusInternalServerError).SendString("Error Internal Server")
	}
}

// TODO: Edit student
func EditStudent(c *fiber.Ctx) error {
	var student struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Grade      string `json:"grade"`
		Attendance bool   `json:"attendance"`
	}

	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	studentJSON, err := json.Marshal(student)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to serialize student data")
	}

	resp, err := sendRequestToCoreValidate("PUT", "/core/edit-user", bytes.NewReader(studentJSON))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error Internal Server")
	}
	defer resp.Body.Close()

	body, err := readResponseBodyValidate(resp)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Failed to read response body")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return c.Status(fiber.StatusOK).Send(body)
	case http.StatusNotFound:
		return c.Status(fiber.StatusNotFound).SendString("Can't edit student")
	default:
		return c.Status(fiber.StatusInternalServerError).SendString("Error Internal Server")
	}
}
