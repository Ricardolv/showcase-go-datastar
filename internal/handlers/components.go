package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ComponentsHandler struct{}

func NewComponentsHandler() *ComponentsHandler {
	return &ComponentsHandler{}
}

func (h *ComponentsHandler) ComponentsPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Components Page - Placeholder",
	})
}
