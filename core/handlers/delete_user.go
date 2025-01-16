package handlers

import (
	"database/sql"
	"fmt"

	"github.com/labstack/gommon/log"
)

func DeleteUser(db *sql.DB, studentId DeleteStudent) error {
	// check student id
	query := "SELECT COUNT(*) FROM students WHERE id=?"

	var count int
	err := db.QueryRow(query, studentId.ID).Scan(&count)
	if count == 0 {
		log.Print("Not Found ID", err)
		return fmt.Errorf("not found student id")
	}

	//delete
	query = "DELETE FROM student_managemrnt.students WHERE id = ?"
	fmt.Println("delete", studentId.ID)

	_, err = db.Exec(query, studentId.ID)
	if err != nil {
		log.Printf("failed to delete student: %v", err)
		return err
	}

	return nil
}
