package main

import (
	"core/handlers"
	"fmt"
	"log"

	"database/sql"

	"github.com/gofiber/fiber/v2"

	_ "github.com/go-sql-driver/mysql"
)

func setupDatabaseConnection() (*sql.DB, error) {

	dsn := "root:@tcp(127.0.0.1:3306)/student_managemrnt"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connection Success")

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func main() {
	app := fiber.New()
	db, err := setupDatabaseConnection()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	//TODO : List
	app.Get("/core/student-list", func(c *fiber.Ctx) error {
		students, err := handlers.StudentList(db)
		if err != nil {
			if err.Error() == "no data" {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "Not have any of student yet",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to fetch student list",
			})
		}
		return c.Status(fiber.StatusOK).JSON(students)
	})

	//TODO : Insert
	app.Post("/core/create-user", func(c *fiber.Ctx) error {
		var student handlers.InsertStudent
		if err := c.BodyParser(&student); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid input data",
			})
		}

		err := handlers.CreateUser(db, student)
		if err != nil {
			if err.Error() == "email existence" {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"message": fmt.Sprintf("This email of %s is already registered. Please use a different email.", student.Name),
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to insert student into the database",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Student created successfully",
		})
	})

	//TODO : Delete
	app.Delete("/core/delete-user", func(c *fiber.Ctx) error {
		var studentId handlers.DeleteStudent

		if err := c.BodyParser(&studentId); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid delete data",
			})
		}

		err := handlers.DeleteUser(db, studentId)
		if err != nil {
			if err.Error() == "not found student id" {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"message": "Not found this student id",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to delete student from the database",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Deleted successfully",
		})
	})

	//TODO :Edit
	app.Put("/core/edit-user", func(c *fiber.Ctx) error {
		var student handlers.EditStudent
		if err := c.BodyParser(&student); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid input data",
			})
		}

		err := handlers.EditUser(db, student)
		if err != nil {
			if err.Error() == "not found student id" {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"message": "Not found this student id",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update student",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Student Edited successfully",
		})
	})

	app.Listen(":3002")
}
