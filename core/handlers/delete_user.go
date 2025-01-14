package handlers

import (
	"database/sql"
	"fmt"

	"github.com/labstack/gommon/log"
)

func DeleteUser(db *sql.DB, studentId DeleteStudent) error {
	query := "DELETE FROM student_managemrnt.students WHERE id = ?"

	fmt.Println("delete", studentId.ID)

	_, err := db.Exec(query, studentId.ID)
	if err != nil {
		log.Fatal("failed to delete student: %v", err)
	}

	return nil
}
