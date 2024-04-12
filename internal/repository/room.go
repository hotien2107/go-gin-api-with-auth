package repository

import (
	"time"

	"gin-rest-api.com/basic/internal/db/postgres"
	"gin-rest-api.com/basic/internal/models"
)

type RoomRepository struct {
	*postgres.PsqlDB
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{
		postgres.NewPsqlDB(),
	}
}

func (r *RoomRepository) GetByUser(userId int64) (*[]models.Room, error) {
	//query string
	query := `
		SELECT * FROM rooms
		WHERE createdUser= $1
	`

	rows, err := r.DB.Query(query, userId)
	if err != nil {
		return &[]models.Room{}, err
	}

	var rooms []models.Room

	for rows.Next() {
		var room models.Room

		err := rows.Scan(&room.ID, &room.Name, &room.CreatedUser, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return &rooms, err
		}
		rooms = append(rooms, room)
	}

	return &rooms, nil

}

func (r *RoomRepository) Create(room *models.Room) (int64, error) {
	//query string
	query := `
		INSERT INTO rooms(name, createdUser, createdAt)
		VALUES ($1,$2, $3)
		RETURNING id
	`

	var roomId int64

	err := r.DB.QueryRow(query, room.Name, room.CreatedUser, time.Now()).Scan(&roomId)
	if err != nil {
		return 0, err
	}

	return roomId, nil
}
