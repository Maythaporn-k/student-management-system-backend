package handlers

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Grade      string `json:"grade"`
	Email      string `json:"email"`
	Attendance bool   `json:"attendance"`
}

type InsertStudent struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

type DeleteStudent struct {
	ID int `json:"id"`
}

type EditStudent struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Grade      string `json:"grade"`
	Attendance bool   `json:"attendance"`
}
