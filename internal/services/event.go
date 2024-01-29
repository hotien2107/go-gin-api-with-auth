package services

import (
	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/repository"
)

type EventService struct {
	repo *repository.EventRepository
}

func NewEventService() *EventService {
	// Initialize and return a new UserService instance
	return &EventService{
		repo: repository.NewEventRepository(),
	}
}

func (e *EventService) GetAll() ([]models.Event, error) {
	events, err := e.repo.GetAll()
	if err != nil {
		return []models.Event{}, err
	}

	return events, nil
}

func (e *EventService) GetById(eventID int64) (models.Event, error) {
	event, err := e.repo.GetById(eventID)
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}

func (e *EventService) Save(event *models.Event) (int64, error) {
	eventId, err := e.repo.Save(event)
	if err != nil {
		return 0, err
	}

	return eventId, nil
}

func (e *EventService) Update(event *models.Event) error {
	err := e.repo.Update(event)
	return err
}

func (e *EventService) Delete(eventId int64) error {
	err := e.repo.Delete(eventId)
	return err
}
