package api

import (
	"gin-rest-api.com/basic/internal/handlers"
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
func (a *API) Start(port string) error {
	a.initializeRoutes()
	return a.engine.Run(":" + port)
}

// initializeRoutes sets up the API routes.
func (a *API) initializeRoutes() {
	eventHandler := handlers.NewEventHandler()
	authHandler := handlers.NewAuthHandler()

	apiAuthV1 := a.engine.Group("/api/v1/")
	{
		apiAuthV1.POST("/sign-up", authHandler.SignUp)
		apiAuthV1.POST("/login", authHandler.Login)
	}

	apiEventV1 := a.engine.Group("/api/v1/event")
	{
		apiEventV1.GET("/get-all", eventHandler.GetAllEvents)
		apiEventV1.GET("/:id", eventHandler.GetEventById)
		apiEventV1.POST("/create", eventHandler.CreateNewEvent)
		apiEventV1.PUT("/:id", eventHandler.UpdateEvent)
		apiEventV1.DELETE("/:id", eventHandler.DeleteEventByID)
	}
}
