package models

import "time"

type Participant struct {
	ID        string    `json:"id"`
	UserId    int64     `json:"userId"`
	RoomId    int64     `json:"roomId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
