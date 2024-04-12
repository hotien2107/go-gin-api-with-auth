package services

import (
	"errors"

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
	isExist := s.repo.CheckParticipantExist(roomId, userId)
	if isExist {
		return errors.New("user has joined room")
	}
	return s.repo.JoinRoom(roomId, userId)
}

func (s *ParticipantService) GetAllParticipantInRoom(roomId int64) (*[]models.UserPublic, error) {

	return s.repo.GetAllParticipantInRoom(roomId)
}
