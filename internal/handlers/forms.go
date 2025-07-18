package handlers

import (
	"encoding/json"
	"net/http"
	"showcase-datastar-go/internal/services"
	"showcase-datastar-go/internal/templates/pages"
	"time"

	"github.com/gin-gonic/gin"
)

type FormsHandler struct {
	formsService *services.FormsService
}

func NewFormsHandler(formsService *services.FormsService) *FormsHandler {
	return &FormsHandler{
		formsService: formsService,
	}
}

// FormsPage renders the forms page
func (h *FormsHandler) FormsPage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	pages.Forms().Render(c.Request.Context(), c.Writer)
}

// ValidateEmail validates email field
func (h *FormsHandler) ValidateEmail(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result := h.formsService.ValidateEmail(req.Email)

	// Update store with validation result
	storeUpdate := map[string]interface{}{}

	if !result.Valid {
		storeUpdate["contactErrors"] = map[string]string{
			"email": result.Message,
		}
		storeUpdate["newsletterError"] = result.Message
	} else {
		storeUpdate["contactErrors"] = map[string]string{}
		storeUpdate["newsletterError"] = ""
	}

	storeData, _ := json.Marshal(storeUpdate)
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "no-cache")
	c.Header("Datastar-Merge-Store", string(storeData))

	c.JSON(http.StatusOK, result)
}

// ValidateField validates generic form fields
func (h *FormsHandler) ValidateField(c *gin.Context) {
	var req struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result := h.formsService.ValidateField(req.Field, req.Value)

	// Update store with validation result
	storeUpdate := map[string]interface{}{}

	if !result.Valid {
		storeUpdate["contactErrors"] = map[string]string{
			req.Field: result.Message,
		}
	} else {
		storeUpdate["contactErrors"] = map[string]string{}
	}

	storeData, _ := json.Marshal(storeUpdate)
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "no-cache")
	c.Header("Datastar-Merge-Store", string(storeData))

	c.JSON(http.StatusOK, result)
}

// ValidateCPF validates Brazilian CPF
func (h *FormsHandler) ValidateCPF(c *gin.Context) {
	var req struct {
		CPF string `json:"cpf"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result := h.formsService.ValidateCPF(req.CPF)

	// Update store with validation result
	storeUpdate := map[string]interface{}{
		"validationResults": map[string]interface{}{
			"cpf": map[string]interface{}{
				"status": func() string {
					if result.Valid {
						return "valid"
					}
					return "invalid"
				}(),
				"message": result.Message,
				"data":    result.Data,
			},
		},
	}

	// Update the CPF field with formatted value if valid
	if result.Valid && result.Data != "" {
		storeUpdate["validation"] = map[string]string{
			"cpf": result.Data,
		}
		// Simplified store update for Datastar
		storeUpdate["validationResults"] = map[string]string{
			"cpf": "valid",
		}
	} else {
		storeUpdate["validationResults"] = map[string]string{
			"cpf": "invalid",
		}
	}

	storeData, _ := json.Marshal(storeUpdate)
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "no-cache")
	c.Header("Datastar-Merge-Store", string(storeData))

	c.JSON(http.StatusOK, result)
}

// MaskPhone applies phone mask
func (h *FormsHandler) MaskPhone(c *gin.Context) {
	var req struct {
		Phone string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	maskedPhone := h.formsService.MaskPhone(req.Phone)

	// Update store with masked value
	storeUpdate := map[string]interface{}{
		"validation": map[string]string{
			"phone": maskedPhone,
		},
	}

	storeData, _ := json.Marshal(storeUpdate)
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "no-cache")
	c.Header("Datastar-Merge-Store", string(storeData))

	c.JSON(http.StatusOK, gin.H{
		"phone":          maskedPhone,
		"formatted":      true,
		"originalLength": len(req.Phone),
		"maskedLength":   len(maskedPhone),
	})
}

// ValidateCEP validates CEP and returns location data
func (h *FormsHandler) ValidateCEP(c *gin.Context) {
	var req struct {
		CEP string `json:"cep"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result := h.formsService.ValidateCEP(req.CEP)

	// Update store with validation result
	storeUpdate := map[string]interface{}{}

	if result.Valid {
		storeUpdate["validationResults"] = map[string]interface{}{
			"cep":     "valid",
			"cepData": result.Data,
		}
	} else {
		storeUpdate["validationResults"] = map[string]string{
			"cep": "invalid",
		}
	}

	storeData, _ := json.Marshal(storeUpdate)
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "no-cache")
	c.Header("Datastar-Merge-Store", string(storeData))

	c.JSON(http.StatusOK, result)
}

// CountChars counts characters and updates progress
func (h *FormsHandler) CountChars(c *gin.Context) {
	var req struct {
		Text  string `json:"text"`
		Field string `json:"field"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	charData := h.formsService.CountCharacters(req.Text, req.Field)

	// Update store based on field
	storeUpdate := map[string]interface{}{}

	if req.Field == "counter" {
		storeUpdate["textArea"] = req.Text
	}

	storeData, _ := json.Marshal(storeUpdate)
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "no-cache")
	c.Header("Datastar-Merge-Store", string(storeData))

	c.JSON(http.StatusOK, charData)
}

// SubmitNewsletter handles newsletter subscription
func (h *FormsHandler) SubmitNewsletter(c *gin.Context) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Simulate loading delay
	time.Sleep(500 * time.Millisecond)

	result := h.formsService.SubmitNewsletter(req.Name, req.Email)

	// Update store with result
	storeUpdate := map[string]interface{}{
		"newsletterLoading": false,
	}

	if result.Valid {
		storeUpdate["newsletterSuccess"] = true
		storeUpdate["newsletterError"] = ""
		storeUpdate["newsletter"] = map[string]string{
			"name":  "",
			"email": "",
		}
	} else {
		storeUpdate["newsletterSuccess"] = false
		storeUpdate["newsletterError"] = result.Message
	}

	storeData, _ := json.Marshal(storeUpdate)
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "no-cache")
	c.Header("Datastar-Merge-Store", string(storeData))

	c.JSON(http.StatusOK, result)
}

// SubmitContact handles contact form submission
func (h *FormsHandler) SubmitContact(c *gin.Context) {
	var req struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Simulate loading delay
	time.Sleep(1 * time.Second)

	result := h.formsService.SubmitContact(req.Name, req.Email, req.Subject, req.Message)

	// Update store with result
	storeUpdate := map[string]interface{}{
		"contactLoading": false,
	}

	if result.Valid {
		storeUpdate["contactSuccess"] = true
		storeUpdate["contactErrors"] = map[string]string{}
		storeUpdate["contactForm"] = map[string]string{
			"name":    "",
			"email":   "",
			"subject": "",
			"message": "",
		}
	} else {
		storeUpdate["contactSuccess"] = false
		storeUpdate["contactErrors"] = map[string]string{
			"general": result.Message,
		}
	}

	storeData, _ := json.Marshal(storeUpdate)
	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "no-cache")
	c.Header("Datastar-Merge-Store", string(storeData))

	c.JSON(http.StatusOK, result)
}

// GetNewsletterSubscribers returns all newsletter subscribers (admin endpoint)
func (h *FormsHandler) GetNewsletterSubscribers(c *gin.Context) {
	subscribers := h.formsService.GetNewsletterSubscribers()
	c.JSON(http.StatusOK, gin.H{
		"subscribers": subscribers,
		"total":       len(subscribers),
	})
}

// GetContactMessages returns all contact messages (admin endpoint)
func (h *FormsHandler) GetContactMessages(c *gin.Context) {
	messages := h.formsService.GetContactMessages()
	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
		"total":    len(messages),
	})
}
