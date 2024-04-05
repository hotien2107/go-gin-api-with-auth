package repository

import (
	"gin-rest-api.com/basic/internal/db/postgres"
	"gin-rest-api.com/basic/internal/models"
)

type FileRepository struct {
	*postgres.PsqlDB
}

func NewFileRepository() *FileRepository {
	return &FileRepository{
		postgres.NewPsqlDB(),
	}
}

func (r *FileRepository) SaveFile(file *models.File) (int64, error) {
	//query string
	query := `
		INSERT INTO files(name, description, url, dateTime, userId)
		VALUES ($1,$2, $3, $4, $5)
		RETURNING id
	`

	var fileId int64

	err := r.DB.QueryRow(query, file.Name, file.Description, file.URL, file.DateTime, file.UserId).Scan(&fileId)
	if err != nil {
		return 0, err
	}

	return fileId, nil
}

func (r *FileRepository) CreateNewTag(tag *models.Tag) error {
	//query string
	query := `
		INSERT INTO tags(name, dateTime, userId)
		VALUES ($1, $2, $3)
	`

	_, err := r.DB.Exec(query, tag.Name, tag.DateTime, tag.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (r *FileRepository) GetTagById(tagId int64) (*models.Tag, error) {
	// query string
	query := `
		SELECT * FROM tags
		WHERE id= $1
	`

	row := r.DB.QueryRow(query, tagId)

	var tag models.Tag
	err := row.Scan(&tag.ID, &tag.Name, &tag.DateTime, &tag.UserId)
	if err != nil {
		return &models.Tag{}, err
	}

	return &tag, nil
}

func (r *FileRepository) CreateNewFileTag(fileId int64, tagId int64) error {
	//query string
	query := `
		INSERT INTO file_tag(fileId, tagId)
		VALUES ($1, $2)
	`

	_, err := r.DB.Exec(query, fileId, tagId)
	if err != nil {
		return err
	}

	return nil
}
