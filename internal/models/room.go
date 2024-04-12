package models

import "time"

type Room struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name" binding:"required"`
	CreatedUser int64      `json:"createdUser"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
