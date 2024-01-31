package services

import (
	"log"
	"mime/multipart"

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

func (s *FileService) Upload(ctx *gin.Context, file *multipart.File, fileName string, tag string) (int64, error) {
	url, err := cloudinary.UploadFile(ctx, file, fileName, tag)
	if err != nil {
		return 0, err
	}
	log.Println("New file url: ", url)

	return 1, nil
}
