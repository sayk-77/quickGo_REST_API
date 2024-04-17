package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	GetById(driverID int) (*models.Feedback, error)
	GetAll() ([]*models.Feedback, error)
	CreateNew(newFeedback *models.Feedback) error
	UpdateStatus(id int) error
}

type FeedbackRepositoryImpl struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) *FeedbackRepositoryImpl {
	return &FeedbackRepositoryImpl{
		db: db,
	}
}

func (fr *FeedbackRepositoryImpl) GetById(feedbackId int) (*models.Feedback, error) {
	var feedback *models.Feedback
	if err := fr.db.First(&feedback, feedbackId).Error; err != nil {
		return nil, err
	}
	return feedback, nil
}

func (fr *FeedbackRepositoryImpl) GetAll() ([]*models.Feedback, error) {
	var allFeedback []*models.Feedback
	if err := fr.db.Find(&allFeedback).Error; err != nil {
		return nil, err
	}
	return allFeedback, nil
}

func (fr *FeedbackRepositoryImpl) CreateNew(newFeedback *models.Feedback) error {
	if err := fr.db.Create(newFeedback).Error; err != nil {
		return err
	}
	return nil
}

func (fr *FeedbackRepositoryImpl) UpdateStatus(id int) error {
	var feedback *models.Feedback
	if err := fr.db.First(&feedback, id).Error; err != nil {
		return err
	}
	feedback.Status = "Обработан"
	if err := fr.db.Save(feedback).Error; err != nil {
		return err
	}

	return nil
}
