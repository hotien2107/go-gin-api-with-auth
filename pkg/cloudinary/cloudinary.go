package cloudinary

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

var Cloudinary *cloudinary.Cloudinary

func Init() {
	cloudName := os.Getenv("CLD_NAME")
	apiKey := os.Getenv("CLD_API_KEY")
	apiSecret := os.Getenv("CLD_API_SECRET")
	// Start by creating a new instance of Cloudinary using CLOUDINARY_URL environment variable.
	// Alternatively you can use cloudinary.NewFromParams() or cloudinary.NewFromURL().
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}

	Cloudinary = cld
}

func UploadFile(ctx *gin.Context, file *multipart.File, fileName string, tag string) (string, error) {
	res, err := Cloudinary.Upload.Upload(ctx, *file, uploader.UploadParams{
		PublicID:       fileName,
		Folder:         tag,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})
	if err != nil {
		return "", err
	}

	return res.SecureURL, nil
}
