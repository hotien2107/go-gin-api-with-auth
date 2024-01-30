package cloudinary

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
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
