package services

import (
	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/repository"
	"github.com/gin-gonic/gin"
)

type MessageService struct {
	repo *repository.MessageRepository
}

func NewMessageService() *MessageService {
	// Initialize and return a new UserService instance
	return &MessageService{
		repo: repository.NewMessageRepository(),
	}
}

func (s *MessageService) GetAllMessages(ctx *gin.Context) ([]models.Message, error) {
	userId := ctx.MustGet("userId").(int64)
	allMessages, err := s.repo.GetAllMessage(ctx, userId)
	if err != nil {
		return []models.Message{}, err
	}

	return allMessages, nil
}
