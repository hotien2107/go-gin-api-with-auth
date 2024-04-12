package services

import (
	"errors"
	"strings"
	"time"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/repository"
	"gin-rest-api.com/basic/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	repo *repository.MessageRepository
}

func NewMessageService() *MessageService {
	return &MessageService{
		repo: repository.NewMessageRepository(),
	}
}

func (s *MessageService) GetAlls(ctx *gin.Context) ([]models.Message, error) {
	userId := ctx.MustGet("userId").(int64)
	allMessages, err := s.repo.GetAll(ctx, userId)
	if err != nil {
		return []models.Message{}, err
	}

	return allMessages, nil
}
func (s *MessageService) Send(ctx *gin.Context, messInfo *models.Message) error {
	// validate input data
	if messInfo.RoomId == 0 {
		return errors.New("no room yet")
	}

	if utils.IsEmpty(messInfo.Content) {
		return errors.New("content must not empty")
	}

	if utils.IsEmpty(messInfo.Type) {
		return errors.New("type must not empty")
	}

	// format data
	messInfo.ID = primitive.NewObjectID()
	messInfo.SenderId = ctx.MustGet("userId").(int64)
	messInfo.Content = strings.TrimSpace(messInfo.Content)
	messInfo.Type = strings.TrimSpace(messInfo.Type)
	messInfo.Status = models.MESSAGE_STATUS["SENT"]
	messInfo.CreatedAt = time.Now()

	err := s.repo.Send(ctx, messInfo)
	if err != nil {
		return err
	}
	return nil
}
