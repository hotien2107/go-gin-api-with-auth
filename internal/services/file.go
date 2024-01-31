package services

import (
	"mime/multipart"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/repository"
	"gin-rest-api.com/basic/pkg/cloudinary"
	"github.com/gin-gonic/gin"
)

type FileService struct {
	repo *repository.FileRepository
}

func NewFileService() *FileService {
	// Initialize and return a new UserService instance
	return &FileService{
		repo: repository.NewFileRepository(),
	}
}

func (s *FileService) Upload(ctx *gin.Context, file *multipart.File, fileInfo *models.File, tagId int64) error {
	// Get tag name from tagId
	tag, err := s.repo.GetTagById(tagId)
	if err != nil {
		return err
	}

	// Store file in cloudinary
	url, err := cloudinary.UploadFile(ctx, file, fileInfo.Name, tag.Name)
	if err != nil {
		return err
	}
	fileInfo.URL = url

	// store file info into files table
	fileId, err := s.repo.SaveFile(fileInfo)
	if err != nil {
		return err
	}

	// store file by tag
	err = s.repo.CreateNewFileTag(fileId, tagId)
	if err != nil {
		return err
	}

	return nil
}

func (s *FileService) CreateNewTag(tag *models.Tag) error {
	err := s.repo.CreateNewTag(tag)
	if err != nil {
		return err
	}

	return nil
}
