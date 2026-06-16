package middleware

import (
	"backend-project/internal/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Logger() fiber.Handler {
	l := logger.GetLogger()
	return func(c *fiber.Ctx) error {
		start := time.Now()
		
		// Continue stack
		err := c.Next()

		duration := time.Since(start)
		
		l.Info("HTTP Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("ip", c.IP()),
			zap.String("request_id", c.GetRespHeader("X-Request-ID")),
		)

		return err
	}
}
