package main

import (
	"log"

	"showcase-datastar-go/internal/handlers"
	"showcase-datastar-go/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize services
	searchService := services.NewSearchService()
	dashboardService := services.NewDashboardService()
	formsService := services.NewFormsService()

	// Initialize handlers
	homeHandler := handlers.NewHomeHandler()
	searchHandler := handlers.NewSearchHandler(searchService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)
	formsHandler := handlers.NewFormsHandler(formsService)
	componentsHandler := handlers.NewComponentsHandler()

	// Setup Gin
	r := gin.Default()

	// Middleware
	r.Use(setupCORS())
	r.Use(setupHeaders())

	// Static files
	r.Static("/static", "./web/static")

	// Routes
	setupRoutes(r, searchHandler, dashboardHandler, formsHandler, homeHandler, componentsHandler)

	// Start server
	log.Println("🚀 CEAR Showcase Go rodando em http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}

func setupRoutes(
	r *gin.Engine,
	searchHandler *handlers.SearchHandler,
	dashboardHandler *handlers.DashboardHandler,
	formsHandler *handlers.FormsHandler,
	homeHandler *handlers.HomeHandler,
	componentsHandler *handlers.ComponentsHandler,
) {
	// Home route
	r.GET("/", homeHandler.HomePage)

	// Search routes
	r.GET("/search", searchHandler.SearchPage)
	r.GET("/search/results", searchHandler.SearchResults)
	r.GET("/search/suggestions", searchHandler.GetSuggestions)
	r.GET("/search/live", searchHandler.LiveSearch)

	// Dashboard routes
	r.GET("/dashboard", dashboardHandler.DashboardPage)
	r.GET("/dashboard/stats", dashboardHandler.GetStats)
	r.GET("/dashboard/live-updates", dashboardHandler.LiveUpdates)
	r.GET("/dashboard/widget/:type", dashboardHandler.GetWidget)

	// Forms routes
	r.GET("/forms", formsHandler.FormsPage)
	r.POST("/forms/validate-email", formsHandler.ValidateEmail)
	r.POST("/forms/validate-field", formsHandler.ValidateField)
	r.POST("/forms/validate-cpf", formsHandler.ValidateCPF)
	r.POST("/forms/mask-phone", formsHandler.MaskPhone)
	r.POST("/forms/validate-cep", formsHandler.ValidateCEP)
	r.POST("/forms/count-chars", formsHandler.CountChars)
	r.POST("/forms/submit-newsletter", formsHandler.SubmitNewsletter)
	r.POST("/forms/contact-submit", formsHandler.SubmitContact)

	// Admin routes (optional - for testing)
	r.GET("/admin/newsletter-subscribers", formsHandler.GetNewsletterSubscribers)
	r.GET("/admin/contact-messages", formsHandler.GetContactMessages)

	// Components routes
	r.GET("/components", componentsHandler.ComponentsPage)
}

func setupCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setupHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Headers necessários para Datastar
		c.Header("Cache-Control", "no-cache")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Next()
	}
}
