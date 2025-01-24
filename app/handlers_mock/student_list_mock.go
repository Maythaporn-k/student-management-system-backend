package handlers_mock

import "core/handlers"

func StudentListMock() []handlers.Student {
	// Mock student data
	students := []handlers.Student{
		{
			ID:         1,
			Name:       "John Doe",
			Age:        20,
			Grade:      "A",
			Email:      "john.doe@example.com",
			Attendance: true,
		},
		{
			ID:         2,
			Name:       "Jane Smith",
			Age:        22,
			Grade:      "B",
			Email:      "jane.smith@example.com",
			Attendance: true,
		},
		{
			ID:         3,
			Name:       "Emily Johnson",
			Age:        21,
			Grade:      "C",
			Email:      "emily.johnson@example.com",
			Attendance: false,
		},
	}

	return students
}
