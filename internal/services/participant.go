package services

import (
	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/repository"
	"gin-rest-api.com/basic/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ParticipantService struct {
	repo *repository.ParticipantRepository
}

func NewParticipantService() *ParticipantService {
	return &ParticipantService{
		repo: repository.NewParticipantRepository(),
	}
}

func (s *ParticipantService) JoinRoom(ctx *gin.Context, roomId int64) error {
	userId := utils.GetUserId(ctx)
	return s.repo.JoinRoom(roomId, userId)
}

func (s *ParticipantService) GetAllParticipantInRoom(roomId int64) (*[]models.UserPublic, error) {
	return s.repo.GetAllParticipantInRoom(roomId)
}
