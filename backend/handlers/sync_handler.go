package handlers

import (
	"media-sequencer/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SyncHandler struct {
	syncService *services.SyncService
}

func NewSyncHandler(syncService *services.SyncService) *SyncHandler {
	return &SyncHandler{
		syncService: syncService,
	}
}

type TriggerSyncRequest struct {
	MediaID         int `json:"media_id" binding:"required"`
	DurationSeconds int `json:"duration_seconds" binding:"required"`
}

func (h *SyncHandler) TriggerSync(c *gin.Context) {
	var req TriggerSyncRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err := h.syncService.TriggerSync(req.MediaID, req.DurationSeconds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sync triggered successfully",
	})
}

func (h *SyncHandler) GetSyncStatus(c *gin.Context) {
	status, err := h.syncService.GetSyncStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get sync status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": status,
	})
}
