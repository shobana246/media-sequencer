package handlers

import (
	"log"
	"media-sequencer/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WindowHandler struct {
	windowService *services.WindowService
}

func NewWindowHandler(windowService *services.WindowService) *WindowHandler {
	return &WindowHandler{
		windowService: windowService,
	}
}

func (h *WindowHandler) GetAllWindows(c *gin.Context) {
	windows, err := h.windowService.GetAllWindows()
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch windows",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": windows,
	})
}
