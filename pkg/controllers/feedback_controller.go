package controllers

import (
	"example.com/go/models"
	"example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type FeedbackController struct {
	feedbackService *service.FeedbackService
}

func NewFeedbackController(app *fiber.App, feedbackService *service.FeedbackService) *FeedbackController {
	feedbackController := &FeedbackController{
		feedbackService: feedbackService,
	}
	app.Get("/feedback/all", feedbackController.GetAll)
	app.Get("/feedback/:id", feedbackController.GetById)
	app.Post("/feedback/new", feedbackController.CreateNew)
	app.Put("/feedback/update/:id", feedbackController.UpdateStatus)

	return feedbackController
}

func (fc *FeedbackController) GetById(c *fiber.Ctx) error {
	feedbackId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	feedback, err := fc.feedbackService.GetById(feedbackId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(feedback)
}

func (fc *FeedbackController) GetAll(c *fiber.Ctx) error {
	feedbackRecord, err := fc.feedbackService.GetAll()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(feedbackRecord)
}

func (fc *FeedbackController) CreateNew(c *fiber.Ctx) error {
	var feedback *models.Feedback
	if err := c.BodyParser(&feedback); err != nil {
		c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	feedback.Status = "Новый"

	if err := fc.feedbackService.CreateNew(feedback); err != nil {
		c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}

func (fc *FeedbackController) UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if err := fc.feedbackService.UpdateStatus(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}
