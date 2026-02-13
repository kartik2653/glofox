package main

import (
	"glofox/internal/booking"
	"glofox/internal/class"
	"glofox/internal/config"
	"glofox/internal/database"
	"glofox/internal/middleware"
	"glofox/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	database.Connect(cfg)

	// Auto migrate tables
	database.DB.AutoMigrate(
		&class.Class{},
	)

	database.DB.AutoMigrate(
		&booking.Booking{},
	)

	database.DB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1
				FROM pg_constraint
				WHERE conname = 'fk_class'
			) THEN
				ALTER TABLE bookings
				ADD CONSTRAINT fk_class
				FOREIGN KEY (class_id)
				REFERENCES classes(id)
				ON DELETE CASCADE;
			END IF;
		END
		$$;
		`)

	app := fiber.New()
	app.Use(middleware.Logger())

	classRepo := class.NewClassRepository(database.DB)
	classService := class.NewClassService(classRepo)
	classHandler := class.NewClassHandler(classService)

	bookingRepo := booking.NewBookingRepository(database.DB)
	bookingService := booking.NewBookingService(bookingRepo, classRepo)
	bookingHandler := booking.NewBookingHandler(bookingService)

	router.Setup(app, classHandler, bookingHandler)

	if err := app.Listen(":3001"); err != nil {
		log.Fatal(err)
	}
}
