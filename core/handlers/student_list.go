package handlers

import (
	"database/sql"
	"log"
)

func StudentList(db *sql.DB) ([]Student, error) {
	query := "SELECT id, name, age, grade, email, attendance FROM students"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Grade, &student.Email, &student.Attendance); err != nil {
			log.Fatal(err)
			return nil, err
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return students, nil
}
