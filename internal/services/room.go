package services

import (
	"errors"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/repository"
	"gin-rest-api.com/basic/pkg/utils"
	"github.com/gin-gonic/gin"
)

type RoomService struct {
	repo *repository.RoomRepository
}

func NewRoomService() *RoomService {
	return &RoomService{
		repo: repository.NewRoomRepository(),
	}
}

func (s *RoomService) GetByUser(ctx *gin.Context) (*[]models.Room, error) {
	userId := utils.GetUserId(ctx)
	rooms, err := s.repo.GetByUser(userId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *RoomService) GetById(roomId int64) (*models.Room, error) {
	room, err := s.repo.GetById(roomId)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (s *RoomService) Create(ctx *gin.Context, room *models.Room) (int64, error) {
	room.CreatedUser = utils.GetUserId(ctx)
	if utils.IsEmpty(room.Name) {
		return 0, errors.New("room's name is not empty")
	}
	roomId, err := s.repo.Create(room)
	if err != nil {
		return 0, err
	}
	return roomId, nil
}
