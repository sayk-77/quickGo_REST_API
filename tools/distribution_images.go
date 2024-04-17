package tools

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func DistributionImages(app *fiber.App) {
	imageDir := "./images"

	app.Get("/:filename", func(c *fiber.Ctx) error {
		filename := c.Params("filename")

		imagePath := fmt.Sprintf("%s/%s", imageDir, filename)

		return c.SendFile(imagePath)
	})
}
