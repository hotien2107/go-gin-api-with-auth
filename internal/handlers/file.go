package handlers

import (
	"net/http"
	"strconv"
	"time"

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
	desc := ctx.Request.FormValue("description")
	tagIdStr := ctx.Request.FormValue("tagId")
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: "Cannot parse tagId: " + err.Error(),
			Result:  nil,
		})
		return
	}

	userLogin := ctx.GetInt64("userId")

	var fileInfo *models.File = &models.File{
		Name:        fileName,
		Description: desc,
		DateTime:    time.Now(),
		UserId:      userLogin,
	}

	err = h.services.Upload(ctx, &file, fileInfo, int64(tagId))
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: "Upload file failed: " + err.Error(),
			Result:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		IsError: false,
		Message: "Upload file success!",
		Result:  nil,
	})
}

func (h *FileHandler) CreateNewTag(ctx *gin.Context) {
	var newTag models.Tag
	err := ctx.ShouldBindJSON(&newTag)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	newTag.UserId = ctx.GetInt64("userId")
	newTag.DateTime = time.Now()

	err = h.services.CreateNewTag(&newTag)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			IsError: true,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		IsError: false,
		Message: "Create tag success",
		Result:  nil,
	})
}
