package routes

import (
	"database/sql"
	"media-sequencer/handlers"
	"media-sequencer/repositories"
	"media-sequencer/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {

	// repositories
	windowRepo := repositories.NewWindowRepository(db)
	playlistRepo := repositories.NewPlaylistRepository(db)
	mediaRepo := repositories.NewMediaRepository(db)
	syncRepo := repositories.NewSyncRepository(db)

	// services
	windowService := services.NewWindowService(windowRepo, playlistRepo, syncRepo, mediaRepo)
	mediaService := services.NewMediaService(mediaRepo, windowRepo, playlistRepo)
	syncService := services.NewSyncService(syncRepo, mediaRepo)

	// handlers
	windowHandler := handlers.NewWindowHandler(windowService)
	mediaHandler := handlers.NewMediaHandler(mediaService)
	syncHandler := handlers.NewSyncHandler(syncService)

	// routes
	api := router.Group("/api")
	{
		api.GET("/windows", windowHandler.GetAllWindows)
		api.POST("/windows/:id/media", mediaHandler.AddMediaToWindow)
		api.POST("/sync", syncHandler.TriggerSync)
		api.GET("/sync/status", syncHandler.GetSyncStatus)
	}
}
