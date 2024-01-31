package repository

type FileRepository struct{}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}
