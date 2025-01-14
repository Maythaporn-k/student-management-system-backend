package handlers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
)

func CreateUser(db *sql.DB, student InsertStudent) error {
	query := fmt.Sprintf("SELECT COUNT(*) FROM students")

	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal("failed to execute count query: %w", err)
	}
	if count == 0 {
		fmt.Print("Reset alter auto-increment")
		query = "ALTER TABLE students AUTO_INCREMENT = 1"

		_, err = db.Exec(query)
		if err != nil {
			return err
		}
	}
	query = "INSERT INTO students (name, email, age ,grade) VALUES (?, ?, ?, ?)"
	parts := strings.Split(student.Name, " ")

	email := fmt.Sprintf("%s.%s@school.com",
		strings.ToLower(parts[0]),
		strings.ToLower(string(parts[1][0])))

	_, err = db.Exec(query, student.Name, email, student.Age, student.Grade)
	if err != nil {
		log.Fatal("failed to insert student: %v", err)
	}

	return nil
}
