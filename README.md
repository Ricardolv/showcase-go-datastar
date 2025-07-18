# 🚀 Showcase Go and Datastar

**Demonstração tecnológica:** Go + Datastar + Tailwind CSS + Templ

Uma aplicação moderna que demonstra as capacidades do **Go 1.21+** combinado com **Datastar** (hypermedia framework), **Tailwind CSS** e **Templ** (templates type-safe) para criar experiências web reativas e performáticas.

---

## ✨ **Principais Funcionalidades**

### 🎨 **Design System CEAR**
- **8 cores principais** com paleta vibrante
- **16+ componentes** reutilizáveis (botões, cards, badges, ícones)
- **3 gradientes** customizados 
- **Shadows** com temática laranja
- **Components Gallery** visual completa

### 🔍 **Active Search** 
- **Busca em tempo real** com algoritmo fuzzy
- **15 tecnologias** mockadas com filtros por categoria
- **Server-Sent Events (SSE)** para updates live
- **Debounce inteligente** (300ms) para performance

### 📊 **Dashboard Real-time**
- **Métricas ao vivo** com Go routines
- **Estatísticas de sistema** (CPU, RAM, rede)
- **Gráficos de atividade** em tempo real
- **Subscriber pattern** com SSE para broadcast

### 📝 **Forms Reativos**
- **Validação brasileira**: CPF, CEP, telefone
- **Máscaras automáticas** durante digitação
- **Contadores de caracteres** com barra de progresso
- **Estados visuais**: loading/success/error
- **Submissão AJAX** sem reload de página

---

## 🛠️ **Stack Tecnológica**

| Tecnologia | Versão | Uso |
|------------|--------|-----|
| **Go** | 1.21+ | Backend, HTTP server, concorrência |
| **Gin** | v1.9+ | Framework web, routing, middleware |
| **Templ** | v0.2+ | Templates type-safe, geração HTML |
| **Datastar** | 1.0 | Hypermedia, reatividade, SSE |
| **Tailwind CSS** | 3.4 | Design system, styling |
| **Air** | v1.40+ | Hot reload desenvolvimento |

---

## 🚀 **Como Executar**

### **1. Pré-requisitos**
```bash
# Go 1.21 ou superior
go version

# Node.js e npm para Tailwind CSS
node --version
npm --version
```

### **2. Clonagem e Setup**
```bash
# Clonar repositório
git clone <repository-url>
cd showcase-datastar-go

# Instalar dependências Go
go mod tidy

# Instalar dependências npm
npm install

# Instalar ferramenta templ
go install github.com/a-h/templ/cmd/templ@latest

# Verificar se templ está no PATH
templ --version
```

### **3. Build e Execução**

#### **🔥 Desenvolvimento (Hot Reload)**
```bash
# Terminal 1: CSS watch
npm run css:watch

# Terminal 2: Templ watch + Air server
make dev
```

#### **🏗️ Build Produção**
```bash
# Build completo
make build

# Executar
./bin/showcase
```

#### **🐳 Docker (Opcional)**
```bash
# Build imagem
docker build -t cear-showcase .

# Executar container
docker run -p 8080:8080 cear-showcase
```

### **4. Acesso**
Abrir navegador em: **http://localhost:8080**

---

## 🌐 **Telas e Funcionalidades**

### **🏠 Home** - `http://localhost:8080/`
- **Landing page** com hero section
- **Overview** das funcionalidades
- **Cards de features** com ícones
- **Design responsivo** mobile-first

### **🔍 Active Search** - `http://localhost:8080/search`
**Funcionalidades:**
- ✅ **Busca fuzzy** em tempo real (300ms debounce)
- ✅ **Filtros por categoria**: Frontend, Backend, DevOps, etc.
- ✅ **Ordenação**: relevância, popularidade, nome
- ✅ **15 tecnologias** com descrições detalhadas
- ✅ **Sugestões rápidas** durante digitação
- ✅ **Loading states** e feedback visual

**Como testar:**
1. Digite `react` → veja matching em tempo real
2. Selecione filtro `Frontend` → resultados filtrados
3. Teste ordenação por popularidade
4. Observe debounce durante digitação rápida

### **📊 Dashboard** - `http://localhost:8080/dashboard`
**Funcionalidades:**
- ✅ **Métricas em tempo real** atualizadas a cada 2 segundos
- ✅ **Server-Sent Events** para push automático
- ✅ **Estatísticas de sistema**: CPU, RAM, rede, storage
- ✅ **Atividades recentes** com timestamps
- ✅ **Gráficos visuais** de progresso

**Como testar:**
1. Observe métricas mudando automaticamente
2. Verifique conexão SSE no DevTools (Network → EventSource)
3. Note atividades sendo adicionadas em tempo real
4. CPU/RAM mostram valores reais do sistema Go

### **📝 Forms** - `http://localhost:8080/forms`
**Funcionalidades:**
- ✅ **Newsletter**: validação de email instantânea
- ✅ **Contato**: 4 campos com validação completa
- ✅ **Validação brasileira**: CPF, CEP, telefone
- ✅ **Máscaras automáticas**: (11) 99999-9999, 00000-000
- ✅ **Contador de caracteres** com warning aos 80%
- ✅ **Estados visuais**: loading/success/error

**Como testar:**
1. **Newsletter**: teste email válido/inválido
2. **CPF**: digite `11144477735` → formatação automática
3. **CEP**: digite `01310100` → busca automática (mock)
4. **Telefone**: digite apenas números → máscara aplicada
5. **Mensagem**: observe contador 0-500 caracteres

### **🎨 Components** - `http://localhost:8080/components`
**Funcionalidades:**
- ✅ **Paleta CEAR completa**: 8 cores × 9 tons cada
- ✅ **Gradientes animados** com hover effects
- ✅ **16+ componentes** interativos
- ✅ **Documentação visual** de uso
- ✅ **Código de exemplo** para cada componente

**Seções:**
1. **Cores**: Primary, Secondary, Success, Warning, Error, Info, Purple, Pink
2. **Gradientes**: CEAR, Tech, Success com animações
3. **Shadows**: Orange, Orange-Large, Default
4. **Botões**: Primary, Secondary, Outline em 3 tamanhos
5. **Badges**: 6 cores × 3 tamanhos + status examples
6. **Cards**: Basic, Info, Warning com hover effects
7. **Feature Cards**: Com ícones e animações
8. **Ícones**: 12 ícones SVG coloridos
9. **Typography**: 7 níveis hierárquicos
10. **Forms**: Inputs, selects, textarea estilizados
11. **Documentação**: Guia de implementação

---

## 🧪 **Testes e APIs**

### **APIs REST Disponíveis**

#### **Search APIs**
```bash
# Busca com query
curl "http://localhost:8080/search/results?q=react&category=frontend&sort=popularity"

# Sugestões
curl "http://localhost:8080/search/suggestions?q=go"

# SSE busca ao vivo
curl -N "http://localhost:8080/search/live"
```

#### **Dashboard APIs**
```bash
# Estatísticas atuais
curl "http://localhost:8080/dashboard/stats" | jq

# SSE métricas ao vivo
curl -N "http://localhost:8080/dashboard/live-updates"

# Widget específico
curl "http://localhost:8080/dashboard/widget/activities" | jq
```

#### **Forms APIs**
```bash
# Validar email
curl -X POST "http://localhost:8080/forms/validate-email" \
  -H "Content-Type: application/json" \
  -d '{"email": "test@exemplo.com"}' | jq

# Validar CPF
curl -X POST "http://localhost:8080/forms/validate-cpf" \
  -H "Content-Type: application/json" \
  -d '{"cpf": "11144477735"}' | jq

# Validar CEP
curl -X POST "http://localhost:8080/forms/validate-cep" \
  -H "Content-Type: application/json" \
  -d '{"cep": "01310100"}' | jq

# Submeter newsletter
curl -X POST "http://localhost:8080/forms/submit-newsletter" \
  -H "Content-Type: application/json" \
  -d '{"name": "João", "email": "joao@exemplo.com"}' | jq
```

#### **Components APIs**
```bash
# Paleta de cores completa
curl "http://localhost:8080/components/colors" | jq '.colors.primary'

# Lista de componentes
curl "http://localhost:8080/components/list" | jq

# Design tokens
curl "http://localhost:8080/components/tokens" | jq '.tokens.spacing'
```

### **Admin Endpoints**
```bash
# Lista inscritos newsletter
curl "http://localhost:8080/admin/newsletter-subscribers" | jq

# Lista mensagens de contato
curl "http://localhost:8080/admin/contact-messages" | jq
```

---

## 📁 **Estrutura do Projeto**

```
showcase-datastar-go/
├── 📄 README.md                    # Este arquivo
├── 📄 Makefile                     # Build automation
├── 📄 go.mod, go.sum               # Dependências Go
├── 📄 package.json                 # Dependências npm
├── 📄 tailwind.config.js           # Config Tailwind + CEAR palette
│
├── 📁 cmd/server/                  # 🚀 Entrypoint aplicação
│   └── main.go                     # Servidor principal + routing
│
├── 📁 internal/
│   ├── 📁 handlers/                # 🎯 HTTP handlers (controllers)
│   │   ├── home.go                 # Handler home page
│   │   ├── search.go               # Handlers active search + SSE
│   │   ├── dashboard.go            # Handlers dashboard + SSE  
│   │   ├── forms.go                # Handlers formulários + validação
│   │   └── components.go           # Handlers components gallery + APIs
│   │
│   ├── 📁 services/                # 🔧 Lógica de negócio
│   │   ├── search.go               # Service fuzzy search + mock data
│   │   ├── dashboard.go            # Service métricas + SSE broadcast
│   │   └── forms.go                # Service validações brasileiras
│   │
│   └── 📁 templates/               # 🎨 Templates Templ (type-safe)
│       ├── 📁 layout/              # Layouts reutilizáveis
│       │   └── main.templ          # Layout principal HTML5
│       ├── 📁 components/          # 🧩 Componentes UI
│       │   ├── navbar.templ        # Navegação responsiva
│       │   ├── footer.templ        # Footer com links
│       │   ├── button.templ        # Botões (3 variantes × 3 tamanhos)
│       │   ├── badge.templ         # Badges (6 cores × 3 tamanhos)
│       │   ├── card.templ          # Cards (3 variantes + StatCard)
│       │   ├── feature_card.templ  # Cards com ícones
│       │   └── icon.templ          # 12 ícones SVG
│       ├── 📁 pages/               # 🌐 Páginas principais
│       │   ├── home_simple.templ   # Home page
│       │   ├── search.templ        # Active Search page
│       │   ├── dashboard.templ     # Dashboard page
│       │   ├── forms.templ         # Forms page
│       │   └── components.templ    # Components Gallery
│       └── 📁 fragments/           # 🔗 Fragmentos reutilizáveis
│           └── search_results.templ # Fragment resultados search
│
├── 📁 web/                         # 🌐 Assets estáticos
│   ├── 📄 input.css                # CSS input Tailwind
│   └── 📁 static/
│       ├── 📁 css/
│       │   └── styles.css          # CSS compilado (auto-gerado)
│       └── 📁 js/
│           └── datastar.js         # Datastar.js local (10.54 KiB)
│
├── 📁 ADRS/                        # 📚 Architecture Decision Records
│   ├── ADR-001-arquitetura-geral-demonstracao.md
│   ├── ADR-002-frontend-datastar-templ.md
│   ├── ADR-003-backend-go-simplificado.md
│   └── ADR-004-dependencias-build-go.md
│
└── 📁 bin/                         # 📦 Binários compilados
    └── showcase                    # Executável principal
```

---

## ⚡ **Performance**

### **Benchmarks**
- **🚀 Startup**: < 1 segundo
- **📦 Binary size**: ~15-20 MB (incluindo templates)  
- **🧠 RAM usage**: < 20 MB em idle
- **⚡ Response time**: < 5ms para páginas estáticas
- **🔄 SSE overhead**: < 1MB/min para dashboard

### **Otimizações Implementadas**
- ✅ **Tailwind CSS** minificado em produção
- ✅ **Templates pré-compilados** com Templ
- ✅ **Static assets** servidos diretamente pelo Gin
- ✅ **Debounce** inteligente nas buscas (300-500ms)
- ✅ **Go routines** para concorrência sem blocking
- ✅ **Mutex locks** para thread-safety nas métricas
- ✅ **Memory pools** para reduzir GC pressure

---

## 🎯 **Casos de Uso**

### **Para Developers Go**
- **Template starter** para projetos web modernos
- **Referência** de integração Go + frontend frameworks
- **Exemplos práticos** de SSE, hot reload, build automation

### **Para UI/UX Designers**
- **Design system** completo e documentado
- **Paleta de cores** harmoniosa e acessível
- **Componentes** reutilizáveis com variações

### **Para Arquitetos de Software**
- **ADRs documentados** com decisões técnicas
- **Patterns** de clean architecture em Go
- **Performance benchmarks** reais

---

## 🔧 **Comandos Úteis**

### **Development**
```bash
# Hot reload completo
make dev

# Apenas watch CSS
npm run css:watch

# Apenas watch templates
templ generate --watch

# Lint e format
make fmt
make lint
```

### **Build**
```bash
# Build desenvolvimento
make build-dev

# Build produção
make build

# Clean artifacts
make clean
```

### **Docker**
```bash
# Build imagem
make docker-build

# Run container
make docker-run

# Stop container
make docker-stop
```

### **Testing**
```bash
# Testes unitários
go test ./...

# Test coverage
go test -cover ./...

# Benchmark
go test -bench=. ./...
```

---

## 📋 **Roadmap Futuro**

### **🎯 Melhorias Técnicas**
- [ ] **WebSockets** além de SSE
- [ ] **Database** real (PostgreSQL)
- [ ] **Authentication** com JWT
- [ ] **Rate limiting** e caching
- [ ] **Monitoring** com Prometheus
- [ ] **Testes E2E** com Playwright

### **🎨 UI/UX Enhancements**  
- [ ] **Dark mode** toggle
- [ ] **Animações** com Framer Motion
- [ ] **PWA** capabilities
- [ ] **Accessibility** WCAG 2.1
- [ ] **Internationalization** i18n

### **🚀 DevOps**
- [ ] **CI/CD** com GitHub Actions  
- [ ] **Deployment** automatizado
- [ ] **Health checks** e metrics
- [ ] **Log aggregation**

---

## 🤝 **Contribuição**

1. **Fork** o repositório
2. **Crie** sua feature branch (`git checkout -b feature/nova-funcionalidade`)
3. **Commit** suas mudanças (`git commit -m 'feat: adiciona nova funcionalidade'`)
4. **Push** para a branch (`git push origin feature/nova-funcionalidade`)
5. **Abra** um Pull Request

### **Padrões de Commit**
- `feat:` nova funcionalidade
- `fix:` correção de bug
- `docs:` documentação
- `style:` formatação
- `refactor:` refatoração
- `test:` testes
- `chore:` manutenção

---

## 📄 **Licença**

Este projeto está sob a licença **MIT**. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.


---

## 🙏 **Agradecimentos**

- **[Datastar](https://data-star.dev)** - Framework hypermedia revolucionário
- **[Templ](https://templ.guide)** - Templates type-safe para Go
- **[Tailwind CSS](https://tailwindcss.com)** - Framework CSS utilitário
- **[Gin](https://gin-gonic.com)** - Framework web minimalista para Go
- **[Air](https://github.com/cosmtrek/air)** - Hot reload tool para Go

---

**⭐ Se este projeto foi útil, por favor considere dar uma estrela no GitHub!** 