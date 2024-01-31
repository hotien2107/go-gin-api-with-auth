package repository

import (
	"gin-rest-api.com/basic/internal/db"
	"gin-rest-api.com/basic/internal/models"
)

type FileRepository struct{}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (*FileRepository) SaveFile(file *models.File) error {
	//query string
	query := `
		INSERT INTO files(name, description, url, dateTime, userId)
		VALUES ($1,$2, $3, $4, $5)
	`

	_, err := db.DB.Exec(query, file.Name, file.Description, file.URL, file.DateTime, file.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (*FileRepository) CreateNewTag(tag *models.Tag) error {
	//query string
	query := `
		INSERT INTO tags(name, dateTime, userId)
		VALUES ($1, $2, $3)
	`

	_, err := db.DB.Exec(query, tag.Name, tag.DateTime, tag.UserId)
	if err != nil {
		return err
	}

	return nil
}
