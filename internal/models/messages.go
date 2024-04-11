package models

type Message struct {
	ID             string `json:"id" bson:"_id"`
	ConversationId int64  `json:"conversationId" bson:"conversation_id" binding:"required"`
	SenderId       int64  `json:"senderId" bson:"sender_id" binding:"required"`
	Content        string `json:"content" bson:"content" binding:"required"`
	Type           string `json:"type" bson:"type"`
	Status         string `json:"status" bson:"status"`
}
