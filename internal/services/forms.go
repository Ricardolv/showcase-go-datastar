package services

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type FormsService struct {
	newsletterSubscribers []NewsletterSubscriber
	contactMessages       []ContactMessage
}

type NewsletterSubscriber struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type ContactMessage struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

type ValidationResult struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

type CEPResponse struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func NewFormsService() *FormsService {
	return &FormsService{
		newsletterSubscribers: []NewsletterSubscriber{},
		contactMessages:       []ContactMessage{},
	}
}

// ValidateEmail validates email format
func (fs *FormsService) ValidateEmail(email string) *ValidationResult {
	if email == "" {
		return &ValidationResult{
			Valid:   false,
			Message: "Email é obrigatório",
		}
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return &ValidationResult{
			Valid:   false,
			Message: "Email inválido",
		}
	}

	return &ValidationResult{
		Valid:   true,
		Message: "Email válido",
	}
}

// ValidateField validates generic fields
func (fs *FormsService) ValidateField(field, value string) *ValidationResult {
	switch field {
	case "name":
		if value == "" {
			return &ValidationResult{
				Valid:   false,
				Message: "Nome é obrigatório",
			}
		}
		if len(value) < 2 {
			return &ValidationResult{
				Valid:   false,
				Message: "Nome deve ter pelo menos 2 caracteres",
			}
		}
		if len(value) > 100 {
			return &ValidationResult{
				Valid:   false,
				Message: "Nome muito longo",
			}
		}
		return &ValidationResult{
			Valid:   true,
			Message: "Nome válido",
		}

	case "subject":
		if value == "" {
			return &ValidationResult{
				Valid:   false,
				Message: "Assunto é obrigatório",
			}
		}
		return &ValidationResult{
			Valid:   true,
			Message: "Assunto válido",
		}

	case "message":
		if value == "" {
			return &ValidationResult{
				Valid:   false,
				Message: "Mensagem é obrigatória",
			}
		}
		if len(value) < 10 {
			return &ValidationResult{
				Valid:   false,
				Message: "Mensagem muito curta (mínimo 10 caracteres)",
			}
		}
		if len(value) > 500 {
			return &ValidationResult{
				Valid:   false,
				Message: "Mensagem muito longa (máximo 500 caracteres)",
			}
		}
		return &ValidationResult{
			Valid:   true,
			Message: "Mensagem válida",
		}

	default:
		return &ValidationResult{
			Valid:   false,
			Message: "Campo não reconhecido",
		}
	}
}

// ValidateCPF validates Brazilian CPF
func (fs *FormsService) ValidateCPF(cpf string) *ValidationResult {
	// Remove formatting
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	cpf = strings.ReplaceAll(cpf, " ", "")

	if len(cpf) != 11 {
		return &ValidationResult{
			Valid:   false,
			Message: "CPF deve ter 11 dígitos",
		}
	}

	// Check if all digits are the same
	allSame := true
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return &ValidationResult{
			Valid:   false,
			Message: "CPF inválido",
		}
	}

	// Validate CPF algorithm
	if !fs.validateCPFDigits(cpf) {
		return &ValidationResult{
			Valid:   false,
			Message: "CPF inválido",
		}
	}

	return &ValidationResult{
		Valid:   true,
		Message: "CPF válido",
		Data:    fs.formatCPF(cpf),
	}
}

// validateCPFDigits implements CPF validation algorithm
func (fs *FormsService) validateCPFDigits(cpf string) bool {
	// First digit validation
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}
	remainder := sum % 11
	firstDigit := 0
	if remainder >= 2 {
		firstDigit = 11 - remainder
	}

	if strconv.Itoa(firstDigit) != string(cpf[9]) {
		return false
	}

	// Second digit validation
	sum = 0
	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (11 - i)
	}
	remainder = sum % 11
	secondDigit := 0
	if remainder >= 2 {
		secondDigit = 11 - remainder
	}

	return strconv.Itoa(secondDigit) == string(cpf[10])
}

// formatCPF formats CPF with dots and dash
func (fs *FormsService) formatCPF(cpf string) string {
	if len(cpf) != 11 {
		return cpf
	}
	return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])
}

// MaskPhone applies phone mask
func (fs *FormsService) MaskPhone(phone string) string {
	// Remove all non-digits
	digits := regexp.MustCompile(`\D`).ReplaceAllString(phone, "")

	if len(digits) == 0 {
		return ""
	}

	// Apply mask based on length
	if len(digits) <= 2 {
		return fmt.Sprintf("(%s", digits)
	} else if len(digits) <= 7 {
		return fmt.Sprintf("(%s) %s", digits[0:2], digits[2:])
	} else if len(digits) <= 11 {
		if len(digits) == 11 {
			return fmt.Sprintf("(%s) %s-%s", digits[0:2], digits[2:7], digits[7:11])
		} else {
			return fmt.Sprintf("(%s) %s-%s", digits[0:2], digits[2:6], digits[6:])
		}
	}

	// Max 11 digits
	return fmt.Sprintf("(%s) %s-%s", digits[0:2], digits[2:7], digits[7:11])
}

// ValidateCEP validates and fetches CEP data
func (fs *FormsService) ValidateCEP(cep string) *ValidationResult {
	// Remove formatting
	cep = strings.ReplaceAll(cep, "-", "")
	cep = strings.ReplaceAll(cep, " ", "")

	// Basic validation
	if len(cep) != 8 {
		return &ValidationResult{
			Valid:   false,
			Message: "CEP deve ter 8 dígitos",
		}
	}

	// Check if all digits
	if !regexp.MustCompile(`^\d{8}$`).MatchString(cep) {
		return &ValidationResult{
			Valid:   false,
			Message: "CEP deve conter apenas números",
		}
	}

	// Mock CEP data (in real app, would call ViaCEP API)
	cepData := fs.getMockCEPData(cep)
	if cepData == nil {
		return &ValidationResult{
			Valid:   false,
			Message: "CEP não encontrado",
		}
	}

	return &ValidationResult{
		Valid:   true,
		Message: "CEP válido",
		Data:    fmt.Sprintf("%s - %s, %s", cepData.Localidade, cepData.Bairro, cepData.UF),
	}
}

// getMockCEPData returns mock CEP data
func (fs *FormsService) getMockCEPData(cep string) *CEPResponse {
	mockData := map[string]*CEPResponse{
		"01310100": {
			CEP:        "01310-100",
			Logradouro: "Avenida Paulista",
			Bairro:     "Bela Vista",
			Localidade: "São Paulo",
			UF:         "SP",
		},
		"20040020": {
			CEP:        "20040-020",
			Logradouro: "Rua da Assembleia",
			Bairro:     "Centro",
			Localidade: "Rio de Janeiro",
			UF:         "RJ",
		},
		"30112000": {
			CEP:        "30112-000",
			Logradouro: "Rua da Bahia",
			Bairro:     "Centro",
			Localidade: "Belo Horizonte",
			UF:         "MG",
		},
		"40070110": {
			CEP:        "40070-110",
			Logradouro: "Rua Chile",
			Bairro:     "Centro",
			Localidade: "Salvador",
			UF:         "BA",
		},
		"80010000": {
			CEP:        "80010-000",
			Logradouro: "Rua XV de Novembro",
			Bairro:     "Centro",
			Localidade: "Curitiba",
			UF:         "PR",
		},
	}

	return mockData[cep]
}

// CountCharacters counts and validates character limits
func (fs *FormsService) CountCharacters(text, field string) map[string]interface{} {
	length := len(text)

	result := map[string]interface{}{
		"length":    length,
		"remaining": 0,
		"maxLength": 100,
		"progress":  0.0,
	}

	switch field {
	case "message":
		result["maxLength"] = 500
		result["remaining"] = 500 - length
		result["progress"] = float64(length) / 500.0 * 100
	case "counter":
		result["maxLength"] = 100
		result["remaining"] = 100 - length
		result["progress"] = float64(length) / 100.0 * 100
	default:
		result["remaining"] = 100 - length
		result["progress"] = float64(length) / 100.0 * 100
	}

	if result["progress"].(float64) > 100 {
		result["progress"] = 100.0
	}

	return result
}

// SubmitNewsletter adds email to newsletter
func (fs *FormsService) SubmitNewsletter(name, email string) *ValidationResult {
	// Validate email
	emailValidation := fs.ValidateEmail(email)
	if !emailValidation.Valid {
		return emailValidation
	}

	// Check if already subscribed
	for _, subscriber := range fs.newsletterSubscribers {
		if subscriber.Email == email {
			return &ValidationResult{
				Valid:   false,
				Message: "Este email já está inscrito",
			}
		}
	}

	// Add subscriber
	subscriber := NewsletterSubscriber{
		ID:        fmt.Sprintf("news_%d", time.Now().UnixNano()),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	fs.newsletterSubscribers = append(fs.newsletterSubscribers, subscriber)

	return &ValidationResult{
		Valid:   true,
		Message: "Inscrito com sucesso!",
		Data:    subscriber.ID,
	}
}

// SubmitContact adds contact message
func (fs *FormsService) SubmitContact(name, email, subject, message string) *ValidationResult {
	// Validate required fields
	if name == "" || email == "" || subject == "" || message == "" {
		return &ValidationResult{
			Valid:   false,
			Message: "Todos os campos são obrigatórios",
		}
	}

	// Validate email
	emailValidation := fs.ValidateEmail(email)
	if !emailValidation.Valid {
		return emailValidation
	}

	// Validate field lengths
	nameValidation := fs.ValidateField("name", name)
	if !nameValidation.Valid {
		return nameValidation
	}

	messageValidation := fs.ValidateField("message", message)
	if !messageValidation.Valid {
		return messageValidation
	}

	// Add contact message
	contact := ContactMessage{
		ID:        fmt.Sprintf("contact_%d", time.Now().UnixNano()),
		Name:      name,
		Email:     email,
		Subject:   subject,
		Message:   message,
		CreatedAt: time.Now(),
	}

	fs.contactMessages = append(fs.contactMessages, contact)

	return &ValidationResult{
		Valid:   true,
		Message: "Mensagem enviada com sucesso!",
		Data:    contact.ID,
	}
}

// GetNewsletterSubscribers returns all newsletter subscribers
func (fs *FormsService) GetNewsletterSubscribers() []NewsletterSubscriber {
	return fs.newsletterSubscribers
}

// GetContactMessages returns all contact messages
func (fs *FormsService) GetContactMessages() []ContactMessage {
	return fs.contactMessages
}
