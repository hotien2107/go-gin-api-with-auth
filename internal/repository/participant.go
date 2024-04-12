package repository

import (
	"gin-rest-api.com/basic/internal/db/postgres"
	"gin-rest-api.com/basic/internal/models"
)

type ParticipantRepository struct {
	*postgres.PsqlDB
}

func NewParticipantRepository() *ParticipantRepository {
	return &ParticipantRepository{
		postgres.NewPsqlDB(),
	}
}

func (r *ParticipantRepository) CheckParticipantExist(roomId int64, userId int64) bool {
	query := `
		SELECT p.id FROM  participants p
		WHERE p.userId = $1 AND p.roomId = $2;
	`

	var id int64
	err := r.DB.QueryRow(query, userId, roomId).Scan(&id)
	if err != nil {
		return false
	}
	if id != 0 {
		return true
	}
	return false
}

func (r *ParticipantRepository) JoinRoom(roomId int64, userId int64) error {
	query := `
		INSERT INTO participants (roomId, userId)
		VALUES ($1, $2)
	`

	_, err := r.DB.Exec(query, roomId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ParticipantRepository) GetAllParticipantInRoom(roomId int64) (*[]models.UserPublic, error) {
	query := `
		SELECT u.id, u.email FROM  users u  
		INNER JOIN participants p
		ON u.id = p.userId
		WHERE p.roomId = $1
	`

	var users []models.UserPublic
	rows, err := r.DB.Query(query, roomId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.UserPublic
		err = rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}
