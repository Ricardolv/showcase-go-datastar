# ADR-SHOWCASE-GO-002: Frontend Hipermídia Showcase - Datastar + Tailwind v4 + Templ

## Status
Aceito

## Self Consistency Check ✅

### Verificação de Consistência Interna:
- ✅ **Datastar + Templ**: CDN Datastar integra perfeitamente com templates Go compilados
- ✅ **Type Safety**: Templ garante compile-time safety alinhado com objetivos
- ✅ **Paleta CEAR**: Tailwind v4 com cores primary/secondary implementa paleta consistentemente
- ✅ **Performance**: Templates compilados + CSS minificado = performance otimizada
- ✅ **Responsividade**: Grid systems e breakpoints atendem critério mobile/desktop

### Conflitos Resolvidos:
- 🔧 **CDN vs Performance**: Datastar 10.54 KiB é aceitável para demonstração
- 🔧 **Templ vs Hot Reload**: `templ generate --watch` alinha com air hot reload
- 🔧 **CSS Build**: Tailwind CLI integra com pipeline Go sem conflitos
- 🔧 **Memory Usage**: Templates compilados reduzem allocations vs text/template

### Decisões Validadas:
- ✅ Datastar (10.54 KiB) é consistente com objetivo de simplicidade
- ✅ Templ type safety alinha com stack Go type-safe
- ✅ Tailwind v4 com paleta customizada atende branding CEAR
- ✅ Component-based architecture reutiliza código eficientemente

## Contexto
O frontend do CEAR Showcase Go precisa **demonstrar de forma impactante** as capacidades do **Datastar** como framework hipermídia integrado com **templ** (Go templates) e a **paleta visual CEAR** implementada com Tailwind CSS v4:

- **Interatividade sem JavaScript complexo**: Demonstrar reatividade pura via hipermídia
- **Paleta visual marcante**: Evidenciar branco, preto e laranjado em componentes reais  
- **Exemplos práticos**: Reproduzir patterns do [data-star.dev](https://data-star.dev/examples/active_search)
- **Type safety**: Aproveitar templ para templates compilados e type-safe
- **Performance Go**: Demonstrar velocidade de renderização server-side
- **Didático e inspirador**: Código que serve de referência para implementação real

## Decisão
Utilizaremos **Datastar** como framework principal com **Tailwind CSS v4** personalizado e **templ** para templates type-safe, criando exemplos **práticos e visuais** que demonstrem todo o potencial da stack Go.

### Tecnologias Escolhidas:
- **Datastar**: Framework hipermídia reativo (10.54 KiB) via CDN
- **Tailwind CSS v4**: Framework CSS utilitário com paleta CEAR
- **Templ**: Engine de templates Go type-safe e compilados
- **CSS Puro**: Animações e transições sem JavaScript adicional

## Arquitetura Frontend Showcase

```
┌─────────────────────────────────────────────────────────────┐
│                    Frontend Architecture                    │
│                                                             │
│  ┌───────────────────┐         ┌───────────────────┐        │
│  │   Templ Files     │         │   Tailwind v4     │        │
│  │   (Compiled)      │◄───────►│   (Build-Time)    │        │
│  │                   │         │                   │        │
│  │ • layout.templ    │         │ • Paleta CEAR     │        │
│  │ • search.templ    │         │ • Componentes     │        │
│  │ • dashboard.templ │         │ • Utilities       │        │
│  │ • forms.templ     │         │ • Animations      │        │
│  └───────────────────┘         └───────────────────┘        │
│           │                             │                   │
│           ▼                             ▼                   │
│  ┌─────────────────────────────────────────────────────┐   │
│  │                  Datastar Layer                     │   │
│  │  • data-on-* (eventos)                              │   │
│  │  • data-bind-* (data binding)                       │   │
│  │  • data-target (DOM updates)                        │   │
│  │  • @get/@post (HTTP requests)                       │   │
│  │  • Server-Sent Events                               │   │
│  └─────────────────────────────────────────────────────┘   │
│           │                                                 │
│           ▼                                                 │
│  ┌─────────────────────────────────────────────────────┐   │
│  │                Go Gin Handlers                      │   │
│  │  • HTML Fragments Response                          │   │
│  │  • JSON Data Response                               │   │
│  │  • SSE Endpoints (Goroutines)                       │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

## Exemplos Templ + Datastar para Showcase

### 1. **Active Search** (baseado no data-star.dev)

```go
// internal/templates/pages/search.templ
package pages

import "cear-showcase/internal/models"

templ SearchPage(title string, description string) {
	@layout.Main(title) {
		<div class="max-w-4xl mx-auto p-6">
			
			<!-- Input de busca com Datastar -->
			<div class="mb-8">
				<input 
					type="text" 
					placeholder="Digite para buscar contatos..."
					class="w-full px-4 py-3 text-lg border-2 border-secondary-300 rounded-lg
					       focus:ring-2 focus:ring-primary-500 focus:border-primary-500
					       transition-all duration-200"
					data-bind-search
					data-on-input__debounce.300ms="@get('/search/results')"
				/>
			</div>
			
			<!-- Resultados dinâmicos -->
			<div id="search-results" 
			     class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
				<!-- Resultados carregados via Datastar -->
			</div>
			
			<!-- Loading state -->
			<div data-show="loading" 
			     class="text-center py-12">
				<div class="inline-block w-8 h-8 border-4 border-primary-500 border-t-transparent rounded-full animate-spin"></div>
				<p class="mt-4 text-secondary-600">Buscando...</p>
			</div>
			
		</div>
	}
}

// Fragment para resultados
templ SearchResults(results []models.SearchResult, query string, total int) {
	for _, result := range results {
		<div class="showcase-card animate-fade-in">
			<div class="flex items-center space-x-4">
				<div class="w-12 h-12 bg-gradient-cear rounded-full flex items-center justify-center text-white font-bold">
					{ result.Initials() }
				</div>
				<div>
					<h3 class="font-semibold text-secondary-900">{ result.FullName }</h3>
					<p class="text-secondary-600">{ result.Email }</p>
					<span class="showcase-badge">{ result.Department }</span>
				</div>
			</div>
		</div>
	}
	if len(results) == 0 && query != "" {
		<div class="text-center py-12">
			<p class="text-secondary-600">Nenhum resultado encontrado para "{ query }"</p>
		</div>
	}
}
```

### 2. **Dashboard Widgets Reativos**

```go
// internal/templates/pages/dashboard.templ  
package pages

import "cear-showcase/internal/models"

templ DashboardPage(data models.DashboardData) {
	@layout.Main("Dashboard Interativo") {
		<div class="grid gap-6 md:grid-cols-2 lg:grid-cols-4">
			
			<!-- Widget de estatística que auto-atualiza -->
			<div class="bg-white rounded-lg shadow-orange-md border border-secondary-100 p-6"
			     data-refresh="@get('/dashboard/stats')"
			     data-interval="5s">
				
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm font-medium text-secondary-600">Usuários Online</p>
						<p class="text-2xl font-bold text-primary-600" 
						   data-text="onlineUsers">{ fmt.Sprintf("%d", data.OnlineUsers) }</p>
					</div>
					<div class="w-12 h-12 bg-primary-100 rounded-lg flex items-center justify-center">
						@components.Icon("users", "w-6 h-6 text-primary-600")
					</div>
				</div>
				
				<!-- Indicador de mudança -->
				<div class="mt-4 flex items-center text-sm"
				     data-class-if-gt="changePercent > 0:text-green-600"
				     data-class-if-lt="changePercent < 0:text-red-600">
					<span data-text="changePercent">{ fmt.Sprintf("%.1f", data.ChangePercent) }</span>% desde ontem
				</div>
				
			</div>
			
			<!-- Mais widgets... -->
			@DashboardWidget(models.WidgetData{
				Type:      "revenue",
				Label:     "Receita Hoje", 
				Value:     data.FormattedRevenue(),
				Trend:     "up",
				Percentage: 75,
			})
			
		</div>
	}
}

// Componente de widget reutilizável
templ DashboardWidget(widget models.WidgetData) {
	<div class="showcase-card">
		<div class="text-center">
			<div class="text-3xl font-bold text-gradient-cear mb-2">
				{ widget.Value }
			</div>
			<div class="text-sm text-secondary-600 mb-4">
				{ widget.Label }
			</div>
			<div class="w-full bg-secondary-200 rounded-full h-2">
				<div class="bg-gradient-cear h-2 rounded-full transition-all duration-500" 
				     style={ fmt.Sprintf("width: %d%%", widget.Percentage) }></div>
			</div>
		</div>
	</div>
}
```

### 3. **Formulários com Validação Reativa**

```go
// internal/templates/pages/forms.templ
package pages

import "cear-showcase/internal/models"

templ FormsPage(form models.ContactFormData) {
	@layout.Main("Formulários Reativos") {
		<form class="bg-white rounded-lg shadow-orange p-8 border border-secondary-100 max-w-2xl mx-auto"
		      data-on-submit="@post('/forms/contact')"
		      data-target="#form-response">
			
			<h2 class="text-2xl font-bold text-secondary-900 mb-6">Contato</h2>
			
			<!-- Campo Email com validação -->
			<div class="mb-6">
				<label class="block text-sm font-medium text-secondary-700 mb-2">
					Email
				</label>
				<input 
					type="email" 
					name="email"
					value={ form.Email }
					class="w-full px-3 py-2 border border-secondary-300 rounded-md
					       focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
					data-on-blur="@post('/forms/validate-email')"
					data-target="#email-feedback"
				/>
				
				<!-- Feedback de validação -->
				<div id="email-feedback" class="mt-2">
					<!-- Feedback carregado via Datastar -->
				</div>
			</div>
			
			<!-- Campo Mensagem -->
			<div class="mb-6">
				<label class="block text-sm font-medium text-secondary-700 mb-2">
					Mensagem
				</label>
				<textarea 
					name="message" 
					rows="4"
					class="w-full px-3 py-2 border border-secondary-300 rounded-md
					       focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
					data-on-input="@post('/forms/count-chars')"
					data-target="#char-count">{ form.Message }</textarea>
				
				<!-- Contador de caracteres -->
				<div id="char-count" class="mt-1 text-sm text-secondary-500">
					{ fmt.Sprintf("%d caracteres", len(form.Message)) }
				</div>
			</div>
			
			<!-- Botão Submit -->
			<button 
				type="submit" 
				class="w-full bg-gradient-cear hover:from-primary-600 hover:to-primary-700 
				       text-white font-medium py-3 px-4 rounded-md transition-all duration-200
				       disabled:opacity-50 disabled:cursor-not-allowed"
				data-loading-text="Enviando..."
				data-loading-class="opacity-50 cursor-not-allowed">
				Enviar Mensagem
			</button>
			
			<!-- Resposta do formulário -->
			<div id="form-response" class="mt-6">
				<!-- Resposta carregada via Datastar -->
			</div>
			
		</form>
	}
}

// Fragment para validação de email
templ EmailValidation(result models.ValidationResult) {
	if result.Valid {
		<div class="flex items-center text-green-600 text-sm">
			@components.Icon("check", "w-4 h-4 mr-1")
			<span>Email válido</span>
		</div>
	} else {
		<div class="flex items-center text-red-600 text-sm">
			@components.Icon("x", "w-4 h-4 mr-1")
			<span>{ result.Message }</span>
		</div>
	}
}
```

## Layout Principal e Componentes

### **Layout Base**

```go
// internal/templates/layout/main.templ
package layout

templ Main(title string) {
	<!DOCTYPE html>
	<html lang="pt-BR">
	<head>
		<meta charset="UTF-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		<title>{ title } - CEAR Showcase Go</title>
		
		<!-- Tailwind CSS -->
		<link href="/static/css/styles.css" rel="stylesheet"/>
		
		<!-- Inter Font -->
		<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap" rel="stylesheet"/>
		
		<!-- Datastar -->
		<script src="/static/js/datastar.js"></script>
	</head>
	<body class="bg-secondary-50 text-secondary-900 antialiased">
		
		<!-- Navigation -->
		@components.Navbar()
		
		<!-- Main Content -->
		<main class="min-h-screen py-8">
			{ children... }
		</main>
		
		<!-- Footer -->
		@components.Footer()
		
	</body>
	</html>
}
```

### **Componentes Reutilizáveis**

```go
// internal/templates/components/navbar.templ
package components

templ Navbar() {
	<nav class="showcase-navbar">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
			<div class="flex justify-between h-16">
				
				<!-- Logo -->
				<div class="flex items-center">
					<a href="/" class="flex items-center space-x-2">
						<div class="w-8 h-8 bg-gradient-cear rounded-lg"></div>
						<span class="text-xl font-bold text-secondary-900">CEAR Go</span>
					</a>
				</div>
				
				<!-- Navigation Links -->
				<div class="flex items-center space-x-8">
					<a href="/" class="text-secondary-700 hover:text-primary-600 px-3 py-2 rounded-md font-medium">
						Home
					</a>
					<a href="/search" class="text-secondary-700 hover:text-primary-600 px-3 py-2 rounded-md font-medium">
						Search
					</a>
					<a href="/dashboard" class="text-secondary-700 hover:text-primary-600 px-3 py-2 rounded-md font-medium">
						Dashboard
					</a>
					<a href="/forms" class="text-secondary-700 hover:text-primary-600 px-3 py-2 rounded-md font-medium">
						Forms
					</a>
					<a href="/components" class="text-secondary-700 hover:text-primary-600 px-3 py-2 rounded-md font-medium">
						Components
					</a>
				</div>
				
			</div>
		</div>
	</nav>
}

// internal/templates/components/button.templ
package components

type ButtonVariant string

const (
	ButtonPrimary   ButtonVariant = "primary"
	ButtonSecondary ButtonVariant = "secondary"
	ButtonOutline   ButtonVariant = "outline"
)

templ Button(text string, variant ButtonVariant, attrs templ.Attributes) {
	<button 
		class={
			"inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md transition-all duration-200",
			templ.KV("bg-primary-500 hover:bg-primary-600 text-white shadow-orange", variant == ButtonPrimary),
			templ.KV("bg-secondary-200 hover:bg-secondary-300 text-secondary-900", variant == ButtonSecondary),
			templ.KV("border-primary-500 text-primary-600 hover:bg-primary-50", variant == ButtonOutline),
		}
		{ attrs... }>
		{ text }
	</button>
}

// internal/templates/components/card.templ
package components

templ Card(title string, children ...templ.Component) {
	<div class="showcase-card">
		if title != "" {
			<h3 class="text-lg font-semibold text-secondary-900 mb-4">{ title }</h3>
		}
		<div>
			for _, child := range children {
				@child
			}
		</div>
	</div>
}
```

## Estrutura de Templates Templ

```
internal/templates/
├── layout/
│   ├── main.templ                    # Layout principal com Datastar
│   ├── navigation.templ              # Navegação entre exemplos
│   └── footer.templ                  # Footer com links
├── pages/
│   ├── home.templ                    # Página inicial do showcase
│   ├── search.templ                  # Active Search demo
│   ├── dashboard.templ               # Dashboard interativo
│   ├── forms.templ                   # Formulários reativos
│   ├── components.templ              # Galeria de componentes
│   └── about.templ                   # Sobre o showcase
├── components/
│   ├── button.templ                  # Botões com variantes
│   ├── card.templ                    # Cards reutilizáveis
│   ├── icon.templ                    # Ícones SVG
│   ├── navbar.templ                  # Navegação principal
│   └── footer.templ                  # Footer
└── fragments/
    ├── search-results.templ          # Fragmento de resultados
    ├── widget-stats.templ            # Fragmento de estatísticas
    └── form-validation.templ         # Fragmento de validação
```

## Vantagens do Templ

### ✅ **Type Safety Completo:**
- **Compile-time**: Erros de template detectados no build
- **Structs nativas**: Integração perfeita com Go types
- **Autocompletion**: IDE support completo
- **Refactoring**: Mudanças propagam automaticamente

### ✅ **Performance Superior:**
- **Templates compilados**: Não há parsing em runtime
- **Memory efficient**: Menos allocations que text/template
- **Hot reload**: templ generate --watch
- **Build otimizado**: Templates como código Go

### ✅ **Developer Experience:**
- **Syntax familiar**: HTML + Go expressions
- **Component-based**: Reutilização natural
- **Error handling**: Stack traces claros
- **IDE integration**: Syntax highlighting

## Critérios de Sucesso

- [ ] **Templ Compilation**: `templ generate` funciona sem erros
- [ ] **Type Safety**: Todos os templates compilam type-safe
- [ ] **Datastar Examples**: Todos os patterns de interação funcionam
- [ ] **Paleta CEAR**: Cores evidenciadas em todos os componentes (primary-500: #f97316)
- [ ] **Responsividade**: Funciona perfeitamente em mobile/desktop
- [ ] **Performance**: Primeira renderização < 100ms (templates compilados)
- [ ] **Hot Reload**: `templ generate --watch` funciona
- [ ] **Component Reuse**: Componentes são reutilizados corretamente
- [ ] **Interatividade**: Todas as ações têm feedback visual instantâneo
- [ ] **Acessibilidade**: Navegação por teclado e screen readers
- [ ] **Build Integration**: CSS + Templates integram com pipeline Go sem conflitos

## Comparação: Templ vs JTE

| Aspecto | Templ (Go) | JTE (Java) |
|---------|------------|------------|
| **Type Safety** | ✅ Compile-time | ✅ Compile-time |
| **Performance** | ✅ Compilado (zero runtime parsing) | ✅ Compilado |
| **Syntax** | HTML + Go expressions | HTML + Java |
| **Hot Reload** | ✅ templ watch | ✅ DevTools |
| **IDE Support** | Bom (melhorando) | Excelente |
| **Learning Curve** | Baixa (Go familiar) | Média |
| **Memory Usage** | Muito baixo (~5MB) | Baixo (~20MB) |
| **Build Speed** | Muito rápido (~2s) | Rápido (~10s) |
| **Ecosystem** | Emergente | Maduro |

---
**Data**: Janeiro 2025  
**Autor**: Equipe CEAR  
**Revisores**: Go Team, Frontend, UX/UI  
**Próxima Revisão**: Março 2025 
**Self Consistency**: ✅ Validado em Janeiro 2025 