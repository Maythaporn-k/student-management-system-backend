package handlers

import (
	"database/sql"
	"fmt"
	"log"
)

func StudentList(db *sql.DB) ([]Student, error) {
	// count total student
	query := "SELECT COUNT(*) FROM students"
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Print("failed to execute count query: %w", err)
	}
	if count == 0 {

		return nil, fmt.Errorf("no data")

	}

	//querry student
	query = "SELECT id, name, age, grade, email, attendance FROM students"
	rows, err := db.Query(query)
	if err != nil {
		log.Print("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Grade, &student.Email, &student.Attendance); err != nil {
			log.Print(err)
			return nil, err
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		log.Print(err)
		return nil, err
	}
	return students, nil
}
