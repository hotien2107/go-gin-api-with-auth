package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MESSAGE_STATUS map[string]string = map[string]string{
	"SENDING": "sending",
	"SENT":    "sent",
}

type Message struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	RoomId    int64              `json:"roomId" bson:"room_id" binding:"required"`
	SenderId  int64              `json:"senderId" bson:"sender_id"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	Type      string             `json:"type" bson:"type"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"createAt" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at"`
}
