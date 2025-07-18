# ADR-SHOWCASE-GO-001: Arquitetura Geral de Demonstração - Go + Templ + Datastar

## Status
Aceito

## Self Consistency Check ✅

### Verificação de Consistência Interna:
- ✅ **Objetivo vs Tecnologias**: Go + Templ + Datastar alinham com simplicidade e performance
- ✅ **Sem BD vs Dados Mock**: Consistente - dados em slices/maps atende demonstração
- ✅ **Performance vs Stack**: Go naturalmente entrega startup < 1s e < 20MB RAM
- ✅ **Didático vs Código**: Handlers simples + structs claros = código educativo
- ✅ **Desenvolvimento vs Produção**: Hot reload (dev) + binary único (prod) são compatíveis

### Conflitos Resolvidos:
- 🔧 **Startup vs Hot Reload**: Startup < 1s refere-se ao binary final, não ao hot reload
- 🔧 **Memory Usage**: Especificado como ~10-20MB para dar margem conforme features
- 🔧 **Templ vs Performance**: Templates compilados garantem performance sem sacrificar type safety

### Decisões Validadas:
- ✅ Gin + Templ + Datastar = Stack coerente para hipermídia reativa
- ✅ Dados mock eliminam complexidade de BD mantendo funcionalidade
- ✅ Go performance nativa suporta todos os requisitos não-funcionais

## Contexto
O Showcase Go precisa demonstrar a integração das tecnologias principais em **Go** de forma **simplificada e didática**:

- Evidenciar capacidades do **Datastar** como framework hipermídia reativo
- Mostrar a **paleta visual CEAR** com Tailwind CSS v4
- Demonstrar **Go** com templates server-side usando **templ**
- Ilustrar **arquitetura simples** sem complexidades desnecessárias
- **Sem banco de dados** (dados mock em memória)
- **Sem autenticação** (foco na demonstração técnica)

## Decisão
Implementaremos uma **aplicação Go minimalista** com foco em **hipermídia reativa** e **experiência visual** usando templ como engine de templates.

### Tecnologias Escolhidas:
- **Go 1.21+**: Linguagem principal
- **Gin Framework**: HTTP router rápido e minimalista
- **Templ**: Engine de templates Go type-safe
- **Datastar**: Framework hipermídia reativo (10.54 KiB)
- **Tailwind CSS v4**: Framework CSS com paleta CEAR personalizada
- **Dados Mock**: Sem dependência de banco de dados
- **Go Modules**: Dependency management

## Arquitetura de Demonstração

```
┌────────────────────────────────────────────────────────────┐
│                  Showcase Go Datastar Application              │
│                                                            │
│  ┌─────────────────┐    ┌─────────────────┐                │
│  │   Handlers      │    │   Templ Files   │                │
│  │   (Gin)         │◄──►│   + Datastar    │                │
│  │                 │    │                 │                │
│  │ • HomeHandler   │    │ • layout.templ  │                │
│  │ • SearchHandler │    │ • search.templ  │                │
│  │ • DashHandler   │    │ • dashboard.templ│                │
│  │ • FormHandler   │    │ • forms.templ   │                │
│  └─────────────────┘    └─────────────────┘                │
│           │                                                │
│           ▼                                                │
│  ┌─────────────────┐    ┌─────────────────┐                │
│  │  Mock Services  │    │   Go Structs    │                │
│  │  (interfaces)   │◄──►│  (Response)     │                │
│  │                 │    │                 │                │
│  │ • MockDataSvc   │    │ • SearchResult  │                │
│  │ • DemoGenerator │    │ • DashboardData │                │
│  │ • StaticContent │    │ • FormResponse  │                │
│  └─────────────────┘    └─────────────────┘                │
│                                                            │
│  ┌─────────────────────────────────────────┐               │
│  │              Tailwind v4 + Datastar    │               │
│  │  • Paleta CEAR (Branco, Preto, Laranja)│               │
│  │  • Componentes Reativos                │               │
│  │  • Interações Hipermídia               │               │
│  └─────────────────────────────────────────┘               │
└────────────────────────────────────────────────────────────┘
```

## Funcionalidades de Demonstração

### 1. **Active Search** (inspirado no data-star.dev)
```
📍 /search
- Input com debounce automático
- Busca dinâmica em lista mock
- Atualização via fragmentos HTML
- Demonstra: data-on-input, data-target
```

### 2. **Dashboard Interativo**
```
📍 /dashboard
- Cards reativos com estatísticas mock
- Widgets que atualizam via SSE simulado
- Gráficos em CSS puro
- Demonstra: data-bind, data-show, intervals
```

### 3. **Formulários Reativos**
```
📍 /forms
- Validação em tempo real
- Feedback visual instantâneo
- Múltiplos tipos de input
- Demonstra: data-on-submit, data-on-blur
```

### 4. **Paleta Visual CEAR**
```
🎨 /components
- Showcase de todos os componentes
- Botões, cards, inputs, alerts
- Demonstração da paleta completa
- Guia de estilo interativo
```

## Estrutura de Diretórios

```
ADR_SHOWCASE_GO/
├── cmd/
│   └── server/
│       └── main.go                    # Entry point da aplicação
├── internal/
│   ├── handlers/
│   │   ├── home.go                   # Handler da página inicial
│   │   ├── search.go                 # Handler do Active Search
│   │   ├── dashboard.go              # Handler do Dashboard
│   │   ├── forms.go                  # Handler dos Formulários
│   │   └── components.go             # Handler dos Componentes
│   ├── models/
│   │   ├── search.go                 # Structs para busca
│   │   ├── dashboard.go              # Structs para dashboard
│   │   └── forms.go                  # Structs para formulários
│   ├── services/
│   │   ├── mock_data.go              # Gerador de dados mock
│   │   ├── search_service.go         # Serviço de busca
│   │   └── dashboard_service.go      # Serviço do dashboard
│   └── templates/
│       ├── layout/
│       │   └── main.templ            # Layout principal
│       ├── pages/
│       │   ├── home.templ            # Página inicial
│       │   ├── search.templ          # Active Search
│       │   ├── dashboard.templ       # Dashboard
│       │   └── forms.templ           # Formulários
│       └── components/
│           ├── navbar.templ          # Navegação
│           ├── card.templ            # Cards reutilizáveis
│           └── button.templ          # Botões estilizados
├── web/
│   ├── static/
│   │   ├── css/
│   │   │   └── styles.css            # CSS gerado pelo Tailwind
│   │   └── js/
│   │       └── datastar.js           # Datastar CDN
│   └── input.css                     # CSS source do Tailwind
├── scripts/
│   ├── dev.sh                        # Script de desenvolvimento
│   ├── build.sh                      # Script de build
│   └── watch.sh                      # Script de watch CSS
├── go.mod                            # Go modules
├── go.sum                            # Dependencies lock
├── tailwind.config.js                # Config Tailwind v4
├── package.json                      # Node.js para Tailwind
└── README.md                         # Documentação
```

## Princípios do Showcase

### ✅ **Foco na Demonstração:**
- **Simplicidade**: Código Go limpo e idiomático
- **Interatividade**: Exemplos práticos de Datastar
- **Visual**: Paleta CEAR bem evidenciada
- **Performance**: Servidor Go rápido e responsivo
- **Educativo**: Código comentado e explicativo

### ❌ **Sem Complexidades Desnecessárias:**
- **Sem banco de dados**: Apenas dados mock em slices/maps
- **Sem autenticação**: Foco na demonstração técnica
- **Sem validações complexas**: Validações simples para demo
- **Sem testes extensivos**: Testes básicos apenas

### 🎯 **Objetivos de Demonstração:**

1. **Datastar Capabilities:**
   - Reatividade sem JavaScript complexo
   - Fragmentos HTML dinâmicos
   - Server-Sent Events nativos do Go
   - Debounce e outros modificadores

2. **Tailwind v4 Integration:**
   - Paleta personalizada CEAR
   - Componentes responsivos
   - Dark mode support
   - Animações CSS puras

3. **Go + Templ Simplicity:**
   - Handlers limpos e focados
   - Templates type-safe com templ
   - Structs bem estruturados
   - Hot reload para desenvolvimento

4. **Go Performance:**
   - Startup rápido (~1 segundo)
   - Baixo uso de memória (~10MB)
   - Concorrência nativa para SSE
   - Build otimizado para produção

## Vantagens da Stack Go

### ✅ **Performance Nativa:**
- **Startup ultra-rápido**: < 1 segundo
- **Baixo uso de memória**: ~10-20MB
- **Concorrência nativa**: Goroutines para SSE
- **Build estático**: Binary único sem dependências

### ✅ **Simplicidade de Desenvolvimento:**
- **Type safety**: templ + Go structs
- **Hot reload**: air + templ watch
- **Dependency management**: Go modules nativo
- **Deploy simples**: Binary único

### ✅ **Ecosistema Moderno:**
- **Gin**: Framework HTTP minimalista
- **Templ**: Templates type-safe modernos
- **Air**: Hot reload robusto
- **Tailwind**: CSS moderno integrado

## Consequências

### Positivas:
- **Performance excepcional**: Go é naturalmente rápido
- **Simplicidade extrema**: Menos código, mais produtividade
- **Deploy trivial**: Binary único auto-executável
- **Concorrência nativa**: SSE e WebSockets sem complexidade
- **Type safety**: templ + Go eliminam erros de template

### Negativas:
- **Ecosistema menor**: Menos libraries que Java/Spring
- **Curva de aprendizado**: Desenvolvedores podem precisar aprender Go
- **Tooling**: IDE support pode ser menor que Java
- **Comunidade**: Menor que Spring Boot para web development

## Critérios de Aceitação

- [ ] **Execução**: `go run cmd/server/main.go` funciona imediatamente
- [ ] **Performance**: Startup < 1 segundo (binary compilado), < 20MB RAM
- [ ] **Navegação**: Todas as páginas demo são acessíveis
- [ ] **Datastar**: Todos os exemplos de interatividade funcionam
- [ ] **Tailwind**: Paleta CEAR é evidenciada em todos os componentes
- [ ] **Templ**: Templates compilam e renderizam type-safe
- [ ] **Hot reload**: Mudanças refletem automaticamente (modo desenvolvimento)
- [ ] **Build**: `go build` gera binary funcional < 20MB
- [ ] **Responsivo**: Funciona em desktop, tablet e mobile
- [ ] **Documentação**: README.md explica como executar
- [ ] **Consistency**: Todas as ADRs são implementadas de forma coerente

## Stack Comparison

| Aspecto | Go + Templ | Java + JTE |
|---------|------------|------------|
| **Startup** | ~1s (compilado) | ~5-10s |
| **Memória** | ~15MB | ~200MB+ |
| **Build** | Binary único | JAR + JVM |
| **Type Safety** | ✅ Templ | ✅ JTE |
| **Hot Reload** | ✅ Air (dev) | ✅ DevTools |
| **Ecosystem** | Menor | Maior |
| **Deploy** | Trivial | Complexo |
| **Learning Curve** | Baixa | Média |

---
**Data**: Janeiro 2025  
**Autor**: Equipe CEAR  
**Revisores**: Go Team, Frontend, Arquitetura  
**Próxima Revisão**: Março 2025 
**Self Consistency**: ✅ Validado em Janeiro 2025 