package booking

import (
	"fmt"
	apimodel "glofox/internal/model/api_model"
	"glofox/pkg"
	"glofox/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	bookingService BookingService
}

func NewBookingHandler(bookingService BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	var booking *apimodel.CreateBookingRequest
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.ErrorResponse{Error: pkg.ErrBadRequest})
	}

	err := h.bookingService.CreateBooking(booking)
	if err != nil {
		fmt.Printf("Error creating booking: %v\n", err)
		switch err {
		case pkg.ErrClassIDRequired,
			pkg.ErrUserIDRequired,
			pkg.ErrInvalidBookingDate,
			pkg.ErrClassNotFound,
			pkg.ErrInvalidDateFormat:

			return c.Status(fiber.StatusBadRequest).
				JSON(response.ErrorResponse{Error: err.Error()})
		}

		return c.Status(fiber.StatusInternalServerError).
			JSON(response.ErrorResponse{Error: pkg.ErrInternalServer})
	}
	return c.Status(fiber.StatusCreated).
		JSON(response.SuccessResponse{
			Message: "booking created successfully",
			Data:    booking,
		})
}

func (h *BookingHandler) ListBookings(c *fiber.Ctx) error {
	query := make(map[string]interface{})

	if userID := c.Query("user_id"); userID != "" {
		query["user_id"] = userID
	}
	if classID := c.Query("class_id"); classID != "" {
		query["class_id"] = classID
	}
	if bookingDate := c.Query("booking_date"); bookingDate != "" {
		query["booking_date"] = bookingDate
	}

	skipStr := c.Query("skip", "0")
	limitStr := c.Query("limit", "10")

	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.ErrorResponse{Error: pkg.ErrBadRequest})
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.ErrorResponse{Error: pkg.ErrBadRequest})
	}
	bookings, total, limit, offset, err := h.bookingService.ListBookings(query, skip, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(response.ErrorResponse{Error: pkg.ErrInternalServer})
	}

	return c.Status(fiber.StatusOK).
		JSON(response.SuccessResponse{
			Data: map[string]interface{}{
				"bookings": bookings,
				"total":    total,
				"limit":    limit,
				"skip":     offset,
			},
			Message: "bookings fetched successfully",
		})
}
