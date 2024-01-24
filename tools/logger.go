package tools

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()

	duration := time.Since(start)

	fmt.Printf("[%s] %s %s - %v\n", c.Method(), c.Path(), c.IP(), duration)

	return err
}
