package router

import (
	"glofox/internal/booking"
	"glofox/internal/class"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, classHandler *class.ClassHandler, bookingHandler *booking.BookingHandler) {

	// Public routes
	noAuthRoutes := app.Group("/api/v1")
	noAuthRoutes.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})
	classRoutes := noAuthRoutes.Group("/classes")
	bookingRoutes := noAuthRoutes.Group("/bookings")
	classRoutes.Post("/", classHandler.CreateClass)
	classRoutes.Get("/", classHandler.ListClasses)
	bookingRoutes.Post("/", bookingHandler.CreateBooking)
	bookingRoutes.Get("/", bookingHandler.ListBookings)

	//accept all unknown routes and return 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Route not found",
		})
	})
}
