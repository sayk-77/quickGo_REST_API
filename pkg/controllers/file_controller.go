package controllers

import (
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type FileController struct {
	fileService service.FileService
}

func NewFileController(app *fiber.App, fileService service.FileService) *FileController {
	fileController := &FileController{
		fileService: fileService,
	}
	app.Post("/upload", fileController.Upload)
	return fileController
}

func (controller *FileController) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("photo")
	if err != nil {
		return err
	}
	imageURL, err := controller.fileService.Upload(file)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"imageURL": imageURL})
}
