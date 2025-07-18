package handlers

import (
	"net/http"
	"showcase-datastar-go/internal/services"
	"showcase-datastar-go/internal/templates/fragments"
	"showcase-datastar-go/internal/templates/pages"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	searchService *services.SearchService
}

func NewSearchHandler(searchService *services.SearchService) *SearchHandler {
	return &SearchHandler{
		searchService: searchService,
	}
}

// SearchPage renders the search page
func (h *SearchHandler) SearchPage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	pages.Search().Render(c.Request.Context(), c.Writer)
}

// SearchResults handles search requests and returns HTML fragments
func (h *SearchHandler) SearchResults(c *gin.Context) {
	// Parse query parameters
	query := c.Query("q")
	category := c.DefaultQuery("category", "all")
	sort := c.DefaultQuery("sort", "relevance")

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	// Perform search
	searchParams := services.SearchParams{
		Query:    query,
		Category: category,
		Sort:     sort,
		Offset:   offset,
		Limit:    limit,
	}

	response := h.searchService.Search(searchParams)

	// Set headers for Datastar
	c.Header("Content-Type", "text/html")
	c.Header("Cache-Control", "no-cache")

	// Update loading state via Datastar
	c.Header("Datastar-Merge-Store", `{"loading": false}`)

	// Render search results fragment
	fragments.SearchResults(response.Results, query, response.TotalResults).Render(c.Request.Context(), c.Writer)
}

// GetSuggestions provides search suggestions (autocomplete)
func (h *SearchHandler) GetSuggestions(c *gin.Context) {
	query := c.Query("q")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		limit = 5
	}

	if query == "" || len(query) < 2 {
		c.JSON(http.StatusOK, gin.H{
			"suggestions": []string{},
		})
		return
	}

	// Get search results for suggestions
	searchParams := services.SearchParams{
		Query:    query,
		Category: "all",
		Sort:     "relevance",
		Offset:   0,
		Limit:    limit,
	}

	response := h.searchService.Search(searchParams)

	// Extract suggestions from results
	suggestions := make([]string, 0, len(response.Results))
	for _, result := range response.Results {
		suggestions = append(suggestions, result.Title)
	}

	c.JSON(http.StatusOK, gin.H{
		"suggestions": suggestions,
		"query":       query,
	})
}

// LiveSearch handles Server-Sent Events for real-time search
func (h *SearchHandler) LiveSearch(c *gin.Context) {
	// Set SSE headers
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Cache-Control")

	// Get query parameters
	query := c.Query("q")
	category := c.DefaultQuery("category", "all")
	sort := c.DefaultQuery("sort", "relevance")

	// Perform search
	searchParams := services.SearchParams{
		Query:    query,
		Category: category,
		Sort:     sort,
		Offset:   0,
		Limit:    10,
	}

	response := h.searchService.Search(searchParams)

	// Send SSE event with search results
	c.Header("Datastar-Merge-Store", `{"loading": false}`)

	// Write SSE event
	c.Writer.Write([]byte("event: search-results\n"))
	c.Writer.Write([]byte("data: "))

	// Render results as HTML and send
	fragments.SearchResults(response.Results, query, response.TotalResults).Render(c.Request.Context(), c.Writer)

	c.Writer.Write([]byte("\n\n"))
	c.Writer.Flush()
}
