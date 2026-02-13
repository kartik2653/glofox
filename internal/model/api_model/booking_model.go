package apimodel

type CreateBookingRequest struct {
	ClassID     uint   `json:"class_id"`
	UserID      uint   `json:"user_id"`
	BookingDate string `json:"booking_date"`
}
