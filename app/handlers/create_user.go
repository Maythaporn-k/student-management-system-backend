package handlers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
)

func CreateUser(db *sql.DB, student InsertStudent) error {
	// count total student
	query := "SELECT COUNT(*) FROM students"
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Print("failed to execute count query: %w", err)
	}
	if count == 0 {
		fmt.Print("Reset alter auto-increment")
		query = "ALTER TABLE students AUTO_INCREMENT = 1"

		_, err = db.Exec(query)
		if err != nil {
			return err
		}
	}

	// check duplicate email
	parts := strings.Split(student.Name, " ")
	email := fmt.Sprintf("%s.%s@school.com",
		strings.ToLower(parts[0]),
		strings.ToLower(string(parts[1][0:2])))

	query = "SELECT COUNT(*) FROM students WHERE email = ?"
	var exists int
	err = db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		log.Printf("failed to check email existence: %v", err)
		return err
	}
	if exists > 0 {
		log.Printf("failed email existence: %v", err)
		return fmt.Errorf("email existence")
	}

	// insert
	query = "INSERT INTO students (name, email, age ,grade) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query, student.Name, email, student.Age, student.Grade)
	if err != nil {
		log.Print("failed to insert student: %v", err)
		return fmt.Errorf("failed to insert student")
	}

	return nil
}
