package api

import (
	"gin-rest-api.com/basic/internal/handlers"
	"gin-rest-api.com/basic/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// API struct holds the Gin engine.
type API struct {
	engine *gin.Engine
}

// NewAPI initializes and returns a new API instance.
func NewAPI() *API {
	return &API{
		engine: gin.Default(),
	}
}

// Start initializes the API and starts the HTTP server.
func (a *API) Start() error {
	a.initializeRoutes()
	return a.engine.Run()
}

func (a *API) initializeRoutes() {
	eventHandler := handlers.NewEventHandler()
	authHandler := handlers.NewAuthHandler()
	fileHandler := handlers.NewFileHandler()

	apiV1 := a.engine.Group("/api/v1/")
	{
		apiV1.POST("/sign-up", authHandler.SignUp)
		apiV1.POST("/login", authHandler.Login)
		apiV1.POST("/gen-new-token", authHandler.GenNewAccessToken)
	}

	apiAuthV1 := apiV1.Group("/")
	apiAuthV1.Use(middlewares.Authenticate)

	apiEventV1 := apiAuthV1.Group("/event")
	{
		apiEventV1.GET("/get-all", eventHandler.GetAllEvents)
		apiEventV1.GET("/:id", eventHandler.GetEventById)
		apiEventV1.POST("/create", eventHandler.CreateNewEvent)
		apiEventV1.PUT("/:id", eventHandler.UpdateEvent)
		apiEventV1.DELETE("/:id", eventHandler.DeleteEventById)
	}

	apiFileV1 := apiAuthV1.Group("/file")
	{
		apiFileV1.POST("/upload", fileHandler.Upload)
		apiFileV1.POST("/add-tag", fileHandler.CreateNewTag)
	}
}
