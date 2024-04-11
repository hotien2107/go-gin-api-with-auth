package repository

import (
	"errors"
	"log"

	"gin-rest-api.com/basic/internal/db/mongodb"
	"gin-rest-api.com/basic/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type MessageRepository struct {
	*mongodb.MongoDB
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		mongodb.NewMongoDB(),
	}
}

func (r *MessageRepository) GetAllMessage(ctx *gin.Context, senderId int64) ([]models.Message, error) {
	var allMess []models.Message
	cur, err := r.DB.Collection("messages").Find(ctx, bson.M{"sender_id": senderId})
	if err != nil {
		return []models.Message{}, errors.New(err.Error())
	}

	for cur.Next(ctx) {
		var mess models.Message
		err := cur.Decode(&mess)
		if err != nil {
			log.Fatal(err)
			return []models.Message{}, errors.New(err.Error())
		}

		allMess = append(allMess, mess)

	}
	return allMess, nil
}
