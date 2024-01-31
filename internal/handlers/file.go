package handlers

import (
	"log"
	"net/http"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/services"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	services *services.FileService
}

func NewFileHandler() *FileHandler {
	return &FileHandler{
		services: services.NewFileService(),
	}
}

func (h *FileHandler) Upload(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(2 << 20) // 32 MB
	if err != nil {
		ctx.JSON(http.StatusRequestURITooLong, models.Response{
			IsError: true,
			Message: "File to large: " + err.Error(),
			Result:  nil,
		})
		return
	}

	// Get the file from the form
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: "Cannot get file from request: " + err.Error(),
			Result:  nil,
		})
		return
	}
	defer file.Close()

	fileName := ctx.Request.FormValue("fileName")
	tagId := ctx.Request.FormValue("tagId")
	tagName := "folder"
	log.Printf("tagId is %v", tagId)
	fileId, err := h.services.Upload(ctx, &file, fileName, tagName)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: "Upload file failed: " + err.Error(),
			Result:  nil,
		})
		return
	}
	// allEvents, err := s.services.GetAll()
	// if err != nil {
	// 	ctx.JSON(http.StatusOK, models.Response{
	// 		IsError: true,
	// 		Message: err.Error(),
	// 		Result:  nil,
	// 	})
	// 	return
	// }
	ctx.JSON(http.StatusOK, models.Response{
		IsError: false,
		Message: "",
		Result:  fileId,
	})
}
