package pkg

import "errors"

var (
	ErrBadRequest        = "bad request"
	ErrInternalServer    = "internal server error"
	ErrInvalidDateFormat = errors.New("invalid date format. expected YYYY-MM-DD")

	ErrClassNameRequired  = errors.New("class name is required")
	ErrStartDateRequired  = errors.New("start date is required")
	ErrEndDateRequired    = errors.New("end date is required")
	ErrInvalidCapacity    = errors.New("capacity must be greater than 0")
	ErrInstructorRequired = errors.New("instructor ID is required")
	ErrInvalidClassName   = errors.New("class name must be between 3 and 100 characters")

	ErrUserIDRequired     = errors.New("user ID is required")
	ErrClassIDRequired    = errors.New("class ID is required")
	ErrClassNotFound      = errors.New("class not found")
	ErrInvalidBookingDate = errors.New("booking date must be within the class start and end dates")
)
