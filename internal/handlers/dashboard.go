package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"showcase-datastar-go/internal/services"
	"showcase-datastar-go/internal/templates/pages"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboardService *services.DashboardService
}

func NewDashboardHandler(dashboardService *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

// DashboardPage renders the dashboard page
func (h *DashboardHandler) DashboardPage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	pages.Dashboard().Render(c.Request.Context(), c.Writer)
}

// GetStats returns current dashboard statistics
func (h *DashboardHandler) GetStats(c *gin.Context) {
	stats := h.dashboardService.GetStats()

	// Set headers for Datastar
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "no-cache")

	// Format timestamp for display
	formattedStats := map[string]interface{}{
		"usersOnline":        stats.UsersOnline,
		"usersChange":        stats.UsersChange,
		"pageViews":          formatNumber(stats.PageViews),
		"viewsChange":        fmt.Sprintf("%.1f", stats.ViewsChange),
		"responseTime":       stats.ResponseTime,
		"responseStatus":     stats.ResponseStatus,
		"memoryUsage":        stats.MemoryUsage,
		"memoryPercent":      stats.MemoryPercent,
		"healthStatus":       stats.HealthStatus,
		"cpuUsage":           stats.CPUUsage,
		"diskUsage":          stats.DiskUsage,
		"networkIO":          stats.NetworkIO,
		"networkIOPercent":   stats.NetworkIOPercent,
		"searchActivity":     stats.SearchActivity,
		"dashboardActivity":  stats.DashboardActivity,
		"formsActivity":      stats.FormsActivity,
		"componentsActivity": stats.ComponentsActivity,
		"lastUpdate":         stats.LastUpdate.Format("15:04:05"),
		"connected":          true,
	}

	// Update store via Datastar header
	storeData, _ := json.Marshal(map[string]interface{}{
		"stats":      formattedStats,
		"connected":  true,
		"lastUpdate": stats.LastUpdate.Format("15:04:05"),
		"error":      nil,
	})

	c.Header("Datastar-Merge-Store", string(storeData))
	c.JSON(http.StatusOK, formattedStats)
}

// LiveUpdates handles Server-Sent Events for real-time dashboard updates
func (h *DashboardHandler) LiveUpdates(c *gin.Context) {
	// Set SSE headers
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Cache-Control")

	// Get client ID (use timestamp as simple ID)
	clientID := fmt.Sprintf("client_%d", time.Now().UnixNano())

	// Subscribe to dashboard updates
	updatesChan := h.dashboardService.Subscribe(clientID)
	defer h.dashboardService.Unsubscribe(clientID)

	// Send initial connection event
	c.Writer.Write([]byte("event: connected\n"))
	c.Writer.Write([]byte("data: {\"message\": \"Connected to dashboard updates\"}\n\n"))
	c.Writer.Flush()

	// Keep connection alive and send updates
	for {
		select {
		case stats, ok := <-updatesChan:
			if !ok {
				return
			}

			// Format stats for frontend
			formattedStats := map[string]interface{}{
				"usersOnline":        stats.UsersOnline,
				"usersChange":        stats.UsersChange,
				"pageViews":          formatNumber(stats.PageViews),
				"viewsChange":        fmt.Sprintf("%.1f", stats.ViewsChange),
				"responseTime":       stats.ResponseTime,
				"responseStatus":     stats.ResponseStatus,
				"memoryUsage":        stats.MemoryUsage,
				"memoryPercent":      stats.MemoryPercent,
				"healthStatus":       stats.HealthStatus,
				"cpuUsage":           stats.CPUUsage,
				"diskUsage":          stats.DiskUsage,
				"networkIO":          stats.NetworkIO,
				"networkIOPercent":   stats.NetworkIOPercent,
				"searchActivity":     stats.SearchActivity,
				"dashboardActivity":  stats.DashboardActivity,
				"formsActivity":      stats.FormsActivity,
				"componentsActivity": stats.ComponentsActivity,
				"lastUpdate":         stats.LastUpdate.Format("15:04:05"),
			}

			// Create store update for Datastar
			storeUpdate := map[string]interface{}{
				"stats":      formattedStats,
				"connected":  true,
				"lastUpdate": stats.LastUpdate.Format("15:04:05"),
				"error":      nil,
			}

			storeJSON, _ := json.Marshal(storeUpdate)

			// Send SSE event with store update
			c.Writer.Write([]byte("event: stats-update\n"))
			c.Writer.Write([]byte(fmt.Sprintf("data: %s\n\n", storeJSON)))
			c.Writer.Flush()

		case <-c.Request.Context().Done():
			// Client disconnected
			return

		case <-time.After(30 * time.Second):
			// Send keepalive ping
			c.Writer.Write([]byte("event: ping\n"))
			c.Writer.Write([]byte("data: {\"type\": \"ping\"}\n\n"))
			c.Writer.Flush()
		}
	}
}

// GetWidget returns a specific widget's data
func (h *DashboardHandler) GetWidget(c *gin.Context) {
	widgetType := c.Param("type")

	switch widgetType {
	case "activity":
		activities := h.dashboardService.GetRecentActivity()
		c.JSON(http.StatusOK, gin.H{
			"activities": activities,
		})

	case "stats":
		stats := h.dashboardService.GetStats()
		c.JSON(http.StatusOK, stats)

	default:
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Widget not found",
		})
	}
}

// Helper functions
func formatNumber(n int) string {
	if n >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(n)/1000000)
	} else if n >= 1000 {
		return fmt.Sprintf("%.1fK", float64(n)/1000)
	}
	return strconv.Itoa(n)
}
