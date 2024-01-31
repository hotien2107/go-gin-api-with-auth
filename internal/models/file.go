package models

import "time"

type File struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	URL         string    `json:"url" binding:"required"`
	DateTime    time.Time `json:"createTime"`
	UserId      int       `json:"userId"`
}

type Tag struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name" binding:"required"`
	DateTime time.Time `json:"createTime"`
	UserId   int       `json:"userId"`
}
