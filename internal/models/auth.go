package models

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserPublic struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
