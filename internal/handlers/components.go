package handlers

import (
	"net/http"
	"showcase-datastar-go/internal/templates/pages"

	"github.com/gin-gonic/gin"
)

type ComponentsHandler struct{}

func NewComponentsHandler() *ComponentsHandler {
	return &ComponentsHandler{}
}

// ComponentsPage renders the components gallery page
func (h *ComponentsHandler) ComponentsPage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	pages.Components().Render(c.Request.Context(), c.Writer)
}

// GetColorPalette returns the CEAR color palette as JSON
func (h *ComponentsHandler) GetColorPalette(c *gin.Context) {
	colors := map[string]interface{}{
		"primary": map[string]string{
			"50":  "#FFF7ED",
			"100": "#FFEDD5",
			"200": "#FED7AA",
			"300": "#FDBA74",
			"400": "#FB923C",
			"500": "#F97316", // Main primary
			"600": "#EA580C",
			"700": "#C2410C",
			"800": "#9A3412",
			"900": "#7C2D12",
		},
		"secondary": map[string]string{
			"50":  "#F9FAFB",
			"100": "#F3F4F6",
			"200": "#E5E7EB",
			"300": "#D1D5DB",
			"400": "#9CA3AF",
			"500": "#6B7280", // Main secondary
			"600": "#4B5563",
			"700": "#374151",
			"800": "#1F2937",
			"900": "#111827",
		},
		"success": map[string]string{
			"50":  "#ECFDF5",
			"100": "#D1FAE5",
			"200": "#A7F3D0",
			"300": "#6EE7B7",
			"400": "#34D399",
			"500": "#10B981", // Main success
			"600": "#059669",
			"700": "#047857",
			"800": "#065F46",
			"900": "#064E3B",
		},
		"warning": map[string]string{
			"50":  "#FFFBEB",
			"100": "#FEF3C7",
			"200": "#FDE68A",
			"300": "#FCD34D",
			"400": "#FBBF24",
			"500": "#F59E0B", // Main warning
			"600": "#D97706",
			"700": "#B45309",
			"800": "#92400E",
			"900": "#78350F",
		},
		"error": map[string]string{
			"50":  "#FEF2F2",
			"100": "#FEE2E2",
			"200": "#FECACA",
			"300": "#FCA5A5",
			"400": "#F87171",
			"500": "#EF4444", // Main error
			"600": "#DC2626",
			"700": "#B91C1C",
			"800": "#991B1B",
			"900": "#7F1D1D",
		},
		"info": map[string]string{
			"50":  "#EFF6FF",
			"100": "#DBEAFE",
			"200": "#BFDBFE",
			"300": "#93C5FD",
			"400": "#60A5FA",
			"500": "#3B82F6", // Main info
			"600": "#2563EB",
			"700": "#1D4ED8",
			"800": "#1E40AF",
			"900": "#1E3A8A",
		},
		"purple": map[string]string{
			"50":  "#F5F3FF",
			"100": "#EDE9FE",
			"200": "#DDD6FE",
			"300": "#C4B5FD",
			"400": "#A78BFA",
			"500": "#8B5CF6", // Main purple
			"600": "#7C3AED",
			"700": "#6D28D9",
			"800": "#5B21B6",
			"900": "#4C1D95",
		},
		"pink": map[string]string{
			"50":  "#FDF2F8",
			"100": "#FCE7F3",
			"200": "#FBCFE8",
			"300": "#F9A8D4",
			"400": "#F472B6",
			"500": "#EC4899", // Main pink
			"600": "#DB2777",
			"700": "#BE185D",
			"800": "#9D174D",
			"900": "#831843",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"colors":    colors,
		"gradients": []string{"bg-gradient-cear", "bg-gradient-to-r"},
		"shadows":   []string{"shadow-orange", "shadow-orange-lg", "shadow-lg"},
	})
}

// GetComponents returns available component types
func (h *ComponentsHandler) GetComponents(c *gin.Context) {
	components := map[string]interface{}{
		"buttons": map[string]interface{}{
			"types":  []string{"primary", "secondary", "outline"},
			"sizes":  []string{"small", "medium", "large"},
			"states": []string{"default", "disabled", "loading"},
		},
		"badges": map[string]interface{}{
			"colors": []string{"primary", "secondary", "success", "warning", "error", "info"},
			"sizes":  []string{"small", "medium", "large"},
		},
		"cards": map[string]interface{}{
			"types": []string{"default", "info", "warning"},
		},
		"icons": []string{
			"home", "search", "trending-up", "check", "heart",
			"menu", "x", "refresh", "alert-circle", "external-link",
			"filter", "star",
		},
		"forms": map[string]interface{}{
			"inputs": []string{"text", "email", "password", "number"},
			"types":  []string{"input", "select", "textarea"},
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"components": components,
		"total":      len(components),
	})
}

// GetDesignTokens returns design system tokens
func (h *ComponentsHandler) GetDesignTokens(c *gin.Context) {
	tokens := map[string]interface{}{
		"spacing": map[string]string{
			"xs": "0.5rem", // 8px
			"sm": "1rem",   // 16px
			"md": "1.5rem", // 24px
			"lg": "2rem",   // 32px
			"xl": "3rem",   // 48px
		},
		"typography": map[string]interface{}{
			"fontFamily": "Inter, system-ui, sans-serif",
			"fontSize": map[string]string{
				"xs":   "0.75rem",
				"sm":   "0.875rem",
				"base": "1rem",
				"lg":   "1.125rem",
				"xl":   "1.25rem",
				"2xl":  "1.5rem",
				"3xl":  "1.875rem",
				"4xl":  "2.25rem",
				"5xl":  "3rem",
			},
			"fontWeight": map[string]string{
				"normal":   "400",
				"medium":   "500",
				"semibold": "600",
				"bold":     "700",
			},
		},
		"borderRadius": map[string]string{
			"none": "0px",
			"sm":   "0.125rem",
			"md":   "0.375rem",
			"lg":   "0.5rem",
			"xl":   "0.75rem",
			"full": "9999px",
		},
		"boxShadow": map[string]string{
			"sm":        "0 1px 2px 0 rgb(0 0 0 / 0.05)",
			"default":   "0 1px 3px 0 rgb(0 0 0 / 0.1)",
			"lg":        "0 10px 15px -3px rgb(0 0 0 / 0.1)",
			"orange":    "0 4px 14px 0 rgb(249 115 22 / 0.39)",
			"orange-lg": "0 25px 50px -12px rgb(249 115 22 / 0.25)",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
		"meta": map[string]interface{}{
			"version":     "1.0.0",
			"description": "CEAR Design System Tokens",
			"lastUpdate":  "2025-01-18",
		},
	})
}
