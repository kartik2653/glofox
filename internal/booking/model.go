package booking

import "glofox/internal/model"

type Booking struct {
	model.BaseModel

	ClassID     uint   `gorm:"not null" json:"class_id"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	BookingDate string `gorm:"not null" json:"booking_date"`
}
