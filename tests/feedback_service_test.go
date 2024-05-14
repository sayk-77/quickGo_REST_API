package tests

import (
	"errors"
	"example.com/go/pkg/service"
	"testing"

	"example.com/go/models"
)

type MockFeedbackRepository struct {
	feedbacks []*models.Feedback
}

func (m *MockFeedbackRepository) GetById(feedbackID int) (*models.Feedback, error) {
	for _, feedback := range m.feedbacks {
		if feedback.ID == uint(feedbackID) {
			return feedback, nil
		}
	}
	return nil, errors.New("feedback not found")
}

func (m *MockFeedbackRepository) GetAll() ([]*models.Feedback, error) {
	return m.feedbacks, nil
}

func (m *MockFeedbackRepository) CreateNew(newFeedback *models.Feedback) error {
	m.feedbacks = append(m.feedbacks, newFeedback)
	return nil
}

func (m *MockFeedbackRepository) UpdateStatus(id int) error {
	for _, feedback := range m.feedbacks {
		if feedback.ID == uint(id) {
			feedback.Status = "Updated"
			return nil
		}
	}
	return errors.New("feedback not found")
}

func TestCreateNewFeedback_Success(t *testing.T) {
	mockRepo := &MockFeedbackRepository{}
	feedbackService := service.NewFeedbackService(mockRepo)

	newFeedback := &models.Feedback{
		Name:   "Nikita",
		Email:  "email@email.com",
		Status: "Pending",
	}

	err := feedbackService.CreateNew(newFeedback)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(mockRepo.feedbacks) != 1 {
		t.Errorf("Expected 1 feedback to be created, got %d", len(mockRepo.feedbacks))
	}

	createdFeedback := mockRepo.feedbacks[0]
	if createdFeedback.ID != newFeedback.ID || createdFeedback.Name != newFeedback.Name || createdFeedback.Status != newFeedback.Status {
		t.Error("Created feedback does not match input")
	}
}

func TestGetFeedbackById_Success(t *testing.T) {
	mockRepo := &MockFeedbackRepository{
		feedbacks: []*models.Feedback{{Name: "Nikita", Email: "email@email.com", Status: "Pending"}},
	}
	feedbackService := service.NewFeedbackService(mockRepo)

	feedbackIndex := 0
	feedback, err := feedbackService.GetById(feedbackIndex)
	if err != nil {
		t.Errorf("Ожидалось отсутствие ошибки, получено: %v", err)
	}
	if feedback == nil {
		t.Error("Ожидался объект обратной связи, получено nil")
	}
}

func TestGetAllFeedbacks_Success(t *testing.T) {
	mockRepo := &MockFeedbackRepository{
		feedbacks: []*models.Feedback{
			{Name: "Nikita", Email: "email@email.com", Status: "Pending"},
			{Name: "Nikita", Email: "email@email.com", Status: "Pending"},
		},
	}
	feedbackService := service.NewFeedbackService(mockRepo)

	allFeedbacks, err := feedbackService.GetAll()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if len(allFeedbacks) != 2 {
		t.Errorf("Expected 2 feedbacks, got %d", len(allFeedbacks))
	}
}
