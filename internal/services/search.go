package services

import (
	"sort"
	"strings"
	"time"

	"showcase-datastar-go/internal/templates/fragments"
)

type SearchService struct {
	mockData []fragments.SearchResult
}

func NewSearchService() *SearchService {
	return &SearchService{
		mockData: getMockSearchData(),
	}
}

// SearchParams represents search parameters
type SearchParams struct {
	Query    string
	Category string
	Sort     string
	Offset   int
	Limit    int
}

// SearchResponse represents search response
type SearchResponse struct {
	Results      []fragments.SearchResult
	TotalResults int
	Query        string
	Duration     time.Duration
}

// Search performs the search with fuzzy matching and filtering
func (s *SearchService) Search(params SearchParams) *SearchResponse {
	startTime := time.Now()

	if params.Limit == 0 {
		params.Limit = 10
	}

	var results []fragments.SearchResult

	// Filter by category
	candidates := s.mockData
	if params.Category != "" && params.Category != "all" {
		var filtered []fragments.SearchResult
		for _, item := range candidates {
			if item.Category == params.Category {
				filtered = append(filtered, item)
			}
		}
		candidates = filtered
	}

	// Search and score
	if params.Query != "" {
		var scored []scoredResult
		query := strings.ToLower(params.Query)

		for _, item := range candidates {
			score := s.calculateScore(item, query)
			if score > 0.1 { // Minimum relevance threshold
				scored = append(scored, scoredResult{
					result: item,
					score:  score,
				})
			}
		}

		// Sort by score
		sort.Slice(scored, func(i, j int) bool {
			return scored[i].score > scored[j].score
		})

		// Extract results
		for _, sr := range scored {
			sr.result.Score = sr.score
			results = append(results, sr.result)
		}
	} else {
		results = candidates
	}

	// Apply sorting
	s.applySorting(results, params.Sort)

	totalResults := len(results)

	// Apply pagination
	if params.Offset >= len(results) {
		results = []fragments.SearchResult{}
	} else {
		end := params.Offset + params.Limit
		if end > len(results) {
			end = len(results)
		}
		results = results[params.Offset:end]
	}

	return &SearchResponse{
		Results:      results,
		TotalResults: totalResults,
		Query:        params.Query,
		Duration:     time.Since(startTime),
	}
}

type scoredResult struct {
	result fragments.SearchResult
	score  float64
}

// calculateScore calculates relevance score for fuzzy matching
func (s *SearchService) calculateScore(item fragments.SearchResult, query string) float64 {
	score := 0.0

	title := strings.ToLower(item.Title)
	description := strings.ToLower(item.Description)

	// Exact match gets highest score
	if strings.Contains(title, query) {
		score += 1.0
	} else if strings.Contains(description, query) {
		score += 0.7
	}

	// Partial word matches
	queryWords := strings.Fields(query)
	for _, word := range queryWords {
		if strings.Contains(title, word) {
			score += 0.6
		}
		if strings.Contains(description, word) {
			score += 0.4
		}

		// Check tags
		for _, tag := range item.Tags {
			if strings.Contains(strings.ToLower(tag), word) {
				score += 0.3
			}
		}
	}

	// Fuzzy matching (simple implementation)
	score += s.fuzzyMatch(title, query) * 0.5
	score += s.fuzzyMatch(description, query) * 0.3

	// Boost by popularity (normalize to 0-0.2 range)
	popularityBoost := float64(item.Popularity) / 10000000.0
	if popularityBoost > 0.2 {
		popularityBoost = 0.2
	}
	score += popularityBoost

	return score
}

// Simple fuzzy matching implementation
func (s *SearchService) fuzzyMatch(text, query string) float64 {
	if len(query) == 0 {
		return 0
	}

	text = strings.ToLower(text)
	query = strings.ToLower(query)

	matches := 0
	queryIndex := 0

	for i := 0; i < len(text) && queryIndex < len(query); i++ {
		if text[i] == query[queryIndex] {
			matches++
			queryIndex++
		}
	}

	return float64(matches) / float64(len(query))
}

// applySorting applies sorting to results
func (s *SearchService) applySorting(results []fragments.SearchResult, sortBy string) {
	switch sortBy {
	case "name":
		sort.Slice(results, func(i, j int) bool {
			return results[i].Title < results[j].Title
		})
	case "popularity":
		sort.Slice(results, func(i, j int) bool {
			return results[i].Popularity > results[j].Popularity
		})
	case "recent":
		sort.Slice(results, func(i, j int) bool {
			return results[i].LastUpdate > results[j].LastUpdate
		})
	case "relevance":
		fallthrough
	default:
		// Results are already sorted by relevance/score
	}
}

// getMockSearchData returns mock search data
func getMockSearchData() []fragments.SearchResult {
	return []fragments.SearchResult{
		{
			ID:          "go",
			Title:       "Go (Golang)",
			Description: "Linguagem de programação open source desenvolvida pelo Google. Simples, rápida e confiável para construir software eficiente.",
			Category:    "languages",
			Tags:        []string{"google", "compiled", "concurrent", "backend", "microservices", "cloud"},
			URL:         "https://golang.org",
			Icon:        "trending-up",
			Popularity:  2500000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "javascript",
			Title:       "JavaScript",
			Description: "Linguagem de programação dinâmica essencial para desenvolvimento web, tanto frontend quanto backend com Node.js.",
			Category:    "languages",
			Tags:        []string{"web", "frontend", "backend", "nodejs", "es6", "typescript"},
			URL:         "https://developer.mozilla.org/en-US/docs/Web/JavaScript",
			Icon:        "trending-up",
			Popularity:  12000000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "python",
			Title:       "Python",
			Description: "Linguagem de programação de alto nível, interpretada e de propósito geral. Ideal para IA, ciência de dados e desenvolvimento web.",
			Category:    "languages",
			Tags:        []string{"ai", "data-science", "django", "flask", "machine-learning"},
			URL:         "https://python.org",
			Icon:        "trending-up",
			Popularity:  8500000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "react",
			Title:       "React",
			Description: "Biblioteca JavaScript para construir interfaces de usuário, especialmente single-page applications. Desenvolvida pelo Facebook.",
			Category:    "frameworks",
			Tags:        []string{"facebook", "frontend", "jsx", "virtual-dom", "spa", "hooks"},
			URL:         "https://reactjs.org",
			Icon:        "trending-up",
			Popularity:  11000000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "gin",
			Title:       "Gin Framework",
			Description: "Framework web HTTP de alta performance para Go. Oferece API rápida e minimalista com suporte a middleware.",
			Category:    "frameworks",
			Tags:        []string{"go", "web", "http", "api", "performance", "middleware"},
			URL:         "https://gin-gonic.com",
			Icon:        "trending-up",
			Popularity:  850000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "docker",
			Title:       "Docker",
			Description: "Plataforma de containerização que permite embalar aplicações em containers leves e portáveis.",
			Category:    "tools",
			Tags:        []string{"containers", "devops", "deployment", "microservices", "kubernetes"},
			URL:         "https://docker.com",
			Icon:        "trending-up",
			Popularity:  4200000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "postgresql",
			Title:       "PostgreSQL",
			Description: "Sistema de gerenciamento de banco de dados relacional open source avançado e confiável.",
			Category:    "databases",
			Tags:        []string{"sql", "relational", "acid", "json", "performance", "enterprise"},
			URL:         "https://postgresql.org",
			Icon:        "trending-up",
			Popularity:  1800000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "mongodb",
			Title:       "MongoDB",
			Description: "Banco de dados NoSQL orientado a documentos, oferece alta performance, disponibilidade e escalabilidade.",
			Category:    "databases",
			Tags:        []string{"nosql", "document", "json", "scalability", "big-data"},
			URL:         "https://mongodb.com",
			Icon:        "trending-up",
			Popularity:  2100000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "vue",
			Title:       "Vue.js",
			Description: "Framework JavaScript progressivo para construir interfaces de usuário. Fácil de aprender e altamente performático.",
			Category:    "frameworks",
			Tags:        []string{"frontend", "spa", "progressive", "components", "reactivity"},
			URL:         "https://vuejs.org",
			Icon:        "trending-up",
			Popularity:  3200000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "rust",
			Title:       "Rust",
			Description: "Linguagem de programação de sistemas que oferece memory safety sem garbage collection e performance excepcional.",
			Category:    "languages",
			Tags:        []string{"systems", "memory-safe", "performance", "mozilla", "webassembly"},
			URL:         "https://rust-lang.org",
			Icon:        "trending-up",
			Popularity:  1500000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "kubernetes",
			Title:       "Kubernetes",
			Description: "Sistema open source para automatizar deployment, escalonamento e gerenciamento de aplicações containerizadas.",
			Category:    "tools",
			Tags:        []string{"orchestration", "containers", "devops", "microservices", "google"},
			URL:         "https://kubernetes.io",
			Icon:        "trending-up",
			Popularity:  1900000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "tailwindcss",
			Title:       "Tailwind CSS",
			Description: "Framework CSS utility-first que permite construir designs customizados rapidamente sem deixar o HTML.",
			Category:    "frameworks",
			Tags:        []string{"css", "utility", "responsive", "design", "frontend"},
			URL:         "https://tailwindcss.com",
			Icon:        "trending-up",
			Popularity:  1400000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "redis",
			Title:       "Redis",
			Description: "Estrutura de dados in-memory open source usada como database, cache e message broker.",
			Category:    "databases",
			Tags:        []string{"in-memory", "cache", "performance", "key-value", "pub-sub"},
			URL:         "https://redis.io",
			Icon:        "trending-up",
			Popularity:  1600000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "nextjs",
			Title:       "Next.js",
			Description: "Framework React para produção que oferece server-side rendering, static generation e muitas outras funcionalidades.",
			Category:    "frameworks",
			Tags:        []string{"react", "ssr", "static", "vercel", "full-stack"},
			URL:         "https://nextjs.org",
			Icon:        "trending-up",
			Popularity:  2800000,
			LastUpdate:  "2024-01",
		},
		{
			ID:          "typescript",
			Title:       "TypeScript",
			Description: "Superset tipado do JavaScript que adiciona definições de tipos estáticos opcionais ao JavaScript.",
			Category:    "languages",
			Tags:        []string{"microsoft", "javascript", "types", "frontend", "backend"},
			URL:         "https://typescriptlang.org",
			Icon:        "trending-up",
			Popularity:  6200000,
			LastUpdate:  "2024-01",
		},
	}
}
