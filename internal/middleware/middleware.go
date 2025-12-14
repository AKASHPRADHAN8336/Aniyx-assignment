package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Log request start
		logger.Info("Request started",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("ip", c.IP()),
			zap.String("user-agent", c.Get("User-Agent")),
		)

		err := c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log request completion
		logger.Info("Request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", latency),
			zap.String("ip", c.IP()),
		)

		return err
	}
}
