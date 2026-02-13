package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		latency := time.Since(start)
		status := c.Response().StatusCode()
		method := c.Method()
		path := c.OriginalURL()

		if err != nil {
			log.Printf(
				"[ERROR] %s %s | %d | %v | %v\n",
				method,
				path,
				status,
				latency,
				err,
			)
			return err
		}

		log.Printf(
			"[INFO] %s %s | %d | %v\n",
			method,
			path,
			status,
			latency,
		)

		return nil
	}
}
