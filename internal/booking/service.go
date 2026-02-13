package booking

import (
	"fmt"
	"glofox/internal/class"
	apimodel "glofox/internal/model/api_model"
	"glofox/pkg"
	"time"
)

type BookingService interface {
	CreateBooking(*apimodel.CreateBookingRequest) error
	ListBookings(query map[string]interface{}, skip int, limit int) ([]*Booking, int, int, int, error)
}

type bookingService struct {
	bookingRepo BookingRepository
	clssRepo    class.ClassRepository
}

func NewBookingService(bookingRepo BookingRepository, clssRepo class.ClassRepository) BookingService {
	return &bookingService{bookingRepo: bookingRepo, clssRepo: clssRepo}
}

func (s *bookingService) CreateBooking(booking *apimodel.CreateBookingRequest) error {
	if booking.ClassID == 0 {
		return pkg.ErrClassIDRequired
	}
	if booking.UserID == 0 {
		return pkg.ErrUserIDRequired
	}
	if booking.BookingDate == "" {
		return pkg.ErrInvalidDateFormat
	}
	layout := "2006-01-02"
	if _, err := time.Parse(layout, booking.BookingDate); err != nil {
		return pkg.ErrInvalidDateFormat
	}
	//check if class id and class exists on the given date
	classQuery := map[string]interface{}{
		"id": booking.ClassID,
	}
	classes, _, _, _, err := s.clssRepo.FindClasses(classQuery, 1, 0)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(classes) == 0 {
		return pkg.ErrClassNotFound
	}
	if booking.BookingDate < classes[0].StartDate.Format(layout) || booking.BookingDate > classes[0].EndDate.Format(layout) {
		return pkg.ErrInvalidBookingDate
	}
	internalBooking := &Booking{
		ClassID:     booking.ClassID,
		UserID:      booking.UserID,
		BookingDate: booking.BookingDate,
	}
	if err := s.bookingRepo.Create(internalBooking); err != nil {
		return err
	}
	return nil
}

func (s *bookingService) ListBookings(query map[string]interface{}, skip int, limit int) ([]*Booking, int, int, int, error) {
	bookings, total, limit, offset, err := s.bookingRepo.FindBookings(query, limit, skip)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	return bookings, total, limit, offset, nil
}
