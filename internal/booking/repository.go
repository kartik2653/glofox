package booking

import "gorm.io/gorm"

type BookingRepository interface {
	Create(booking *Booking) error
	FindBookings(query map[string]interface{}, limit int, offset int) ([]*Booking, int, int, int, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(booking *Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) FindBookings(
	query map[string]interface{},
	limit int,
	offset int,
) ([]*Booking, int, int, int, error) {

	var bookings []*Booking

	err := r.db.
		Where(query).
		Limit(limit).
		Offset(offset).
		Find(&bookings).Error

	if err != nil {
		return nil, 0, 0, 0, err
	}

	var total int64

	if err := r.db.
		Model(&Booking{}).
		Where(query).
		Count(&total).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	return bookings, int(total), limit, offset, nil
}
