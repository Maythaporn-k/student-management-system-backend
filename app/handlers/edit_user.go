package handlers

import (
	"database/sql"
	"fmt"
	"log"
)

func EditUser(db *sql.DB, student EditStudent) error {
	// check studentid
	query := "SELECT COUNT(*) FROM students WHERE id=?"
	var count int
	err := db.QueryRow(query, student.ID).Scan(&count)
	if count == 0 {
		log.Print("Not Found ID", err)
		return fmt.Errorf("not found student id")

	}

	//edit
	query = "UPDATE students SET name = ?, age = ?, grade = ?, attendance = ? WHERE id = ?"
	// Execute the update query
	_, err = db.Exec(query, student.Name, student.Age, student.Grade, student.Attendance, student.ID)
	if err != nil {
		log.Printf("failed to update student: %v", err)
		return err
	}

	return nil
}
