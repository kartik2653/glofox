package class

import (
	"glofox/internal/model"
	"time"
)

type Class struct {
	model.BaseModel

	ClassName    string    `gorm:"type:varchar(255);not null" json:"class_name"`
	StartDate    time.Time `gorm:"not null" json:"start_date"`
	EndDate      time.Time `gorm:"not null" json:"end_date"`
	Capacity     int       `gorm:"not null" json:"capacity"`
	InstructorID uint      `gorm:"not null" json:"instructor_id"`
}
