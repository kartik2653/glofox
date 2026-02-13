package apimodel

type Class struct {
	ClassName    string `json:"class_name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Capacity     int    `json:"capacity"`
	InstructorID uint   `json:"instructor_id"`
}

