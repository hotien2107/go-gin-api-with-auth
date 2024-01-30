package api

import (
	"gin-rest-api.com/basic/cmd/docs"
	"gin-rest-api.com/basic/internal/handlers"
	"gin-rest-api.com/basic/internal/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// swagger api
	docs.SwaggerInfo.BasePath = "/api/v1"
	a.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiAuthV1 := a.engine.Group("/api/v1/")
	{
		apiAuthV1.POST("/sign-up", authHandler.SignUp)
		apiAuthV1.POST("/login", authHandler.Login)
	}

	apiEventV1 := a.engine.Group("/api/v1/event")
	apiEventV1Auth := apiEventV1.Group("/")
	apiEventV1Auth.Use(middlewares.Authenticate)
	{
		apiEventV1Auth.GET("/get-all", eventHandler.GetAllEvents)
		apiEventV1Auth.GET("/:id", eventHandler.GetEventById)
		apiEventV1Auth.POST("/create", eventHandler.CreateNewEvent)
		apiEventV1Auth.PUT("/:id", eventHandler.UpdateEvent)
		apiEventV1Auth.DELETE("/:id", eventHandler.DeleteEventByID)
	}
}
