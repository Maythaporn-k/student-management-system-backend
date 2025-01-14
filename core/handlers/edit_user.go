package handlers

import (
	"database/sql"
	"fmt"
	"log"
)

func EditUser(db *sql.DB, student EditStudent) error {
	fmt.Print(student)
	query := "UPDATE students SET name = ?, age = ?, grade = ?, attendance = ? WHERE id = ?"
	// Execute the update query
	_, err := db.Exec(query, student.Name, student.Age, student.Grade, student.Attendance, student.ID)
	if err != nil {
		log.Fatal("failed to update student: %v", err)
	}

	return nil
}
