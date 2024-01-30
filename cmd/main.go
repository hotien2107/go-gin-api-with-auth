package main

import (
	"gin-rest-api.com/basic/internal/api"
	"gin-rest-api.com/basic/internal/db"
	"gin-rest-api.com/basic/pkg/cloudinary"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	db.InitDB()
	cloudinary.Init()
	api := api.NewAPI()

	api.Start()

}
