package handlers

import (
	"media-sequencer/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	mediaService *services.MediaService
}

func NewMediaHandler(mediaService *services.MediaService) *MediaHandler {
	return &MediaHandler{
		mediaService: mediaService,
	}
}

type AddMediaRequest struct {
	Name            string `json:"name" binding:"required"`
	Type            string `json:"type" binding:"required"`
	URL             string `json:"url"`
	DurationSeconds int    `json:"duration_seconds" binding:"required"`
}

func (h *MediaHandler) AddMediaToWindow(c *gin.Context) {
	// get window id from URL
	windowID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid window ID",
		})
		return
	}

	// parse request body
	var req AddMediaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// call service
	err = h.mediaService.AddMediaToWindow(windowID, req.Name, req.Type, req.URL, req.DurationSeconds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add media",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Media added successfully",
	})
}
