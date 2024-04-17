package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
)

type FeedbackService struct {
	feedbackRepository database.FeedbackRepository
}

func NewFeedbackService(feedbackRepository database.FeedbackRepository) *FeedbackService {
	return &FeedbackService{
		feedbackRepository: feedbackRepository,
	}
}

func (fs *FeedbackService) GetById(feedbackId int) (*models.Feedback, error) {
	return fs.feedbackRepository.GetById(feedbackId)
}

func (fs *FeedbackService) GetAll() ([]*models.Feedback, error) {
	return fs.feedbackRepository.GetAll()
}

func (fs *FeedbackService) CreateNew(newFeedback *models.Feedback) error {
	return fs.feedbackRepository.CreateNew(newFeedback)
}

func (fs *FeedbackService) UpdateStatus(id int) error {
	return fs.feedbackRepository.UpdateStatus(id)
}
