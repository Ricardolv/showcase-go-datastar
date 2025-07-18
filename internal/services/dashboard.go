package services

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type DashboardService struct {
	stats        *DashboardStats
	mu           sync.RWMutex
	subscribers  map[string]chan *DashboardStats
	subMu        sync.RWMutex
	stopChan     chan struct{}
	activities   []ActivityItem
	activitiesMu sync.RWMutex
}

type DashboardStats struct {
	// User Metrics
	UsersOnline int     `json:"usersOnline"`
	UsersChange int     `json:"usersChange"`
	PageViews   int     `json:"pageViews"`
	ViewsChange float64 `json:"viewsChange"`

	// Server Metrics
	ResponseTime   string  `json:"responseTime"`
	ResponseStatus string  `json:"responseStatus"`
	MemoryUsage    string  `json:"memoryUsage"`
	MemoryPercent  float64 `json:"memoryPercent"`

	// System Health
	HealthStatus     string  `json:"healthStatus"`
	CPUUsage         float64 `json:"cpuUsage"`
	DiskUsage        float64 `json:"diskUsage"`
	NetworkIO        int     `json:"networkIO"`
	NetworkIOPercent float64 `json:"networkIOPercent"`

	// Activity Metrics
	SearchActivity     float64 `json:"searchActivity"`
	DashboardActivity  float64 `json:"dashboardActivity"`
	FormsActivity      float64 `json:"formsActivity"`
	ComponentsActivity float64 `json:"componentsActivity"`

	// Metadata
	LastUpdate time.Time `json:"lastUpdate"`
}

type ActivityItem struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
}

func NewDashboardService() *DashboardService {
	ds := &DashboardService{
		stats: &DashboardStats{
			UsersOnline:        rand.Intn(50) + 20,
			UsersChange:        rand.Intn(20) + 5,
			PageViews:          rand.Intn(10000) + 5000,
			ViewsChange:        float64(rand.Intn(30)) + 10.5,
			ResponseTime:       "45ms",
			ResponseStatus:     "Excellent",
			MemoryUsage:        "184MB",
			MemoryPercent:      float64(rand.Intn(40)) + 30,
			HealthStatus:       "healthy",
			CPUUsage:           float64(rand.Intn(60)) + 20,
			DiskUsage:          float64(rand.Intn(50)) + 25,
			NetworkIO:          rand.Intn(50) + 10,
			NetworkIOPercent:   float64(rand.Intn(40)) + 10,
			SearchActivity:     float64(rand.Intn(50)) + 25,
			DashboardActivity:  float64(rand.Intn(40)) + 40,
			FormsActivity:      float64(rand.Intn(40)) + 10,
			ComponentsActivity: float64(rand.Intn(30)) + 50,
			LastUpdate:         time.Now(),
		},
		subscribers: make(map[string]chan *DashboardStats),
		stopChan:    make(chan struct{}),
		activities:  []ActivityItem{},
	}

	// Start background goroutine for real-time updates
	go ds.startRealTimeUpdates()

	return ds
}

// GetStats returns current dashboard statistics
func (ds *DashboardService) GetStats() *DashboardStats {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	// Update with real memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	stats := *ds.stats
	stats.MemoryUsage = formatBytes(m.Alloc)
	stats.MemoryPercent = float64(m.Alloc) / float64(m.Sys) * 100
	if stats.MemoryPercent > 100 {
		stats.MemoryPercent = float64(rand.Intn(40)) + 30
	}
	stats.LastUpdate = time.Now()

	return &stats
}

// GetRecentActivity returns recent activities
func (ds *DashboardService) GetRecentActivity() []ActivityItem {
	ds.activitiesMu.RLock()
	defer ds.activitiesMu.RUnlock()

	// Return last 10 activities
	if len(ds.activities) > 10 {
		return ds.activities[len(ds.activities)-10:]
	}
	return ds.activities
}

// Subscribe adds a new subscriber for real-time updates
func (ds *DashboardService) Subscribe(id string) chan *DashboardStats {
	ds.subMu.Lock()
	defer ds.subMu.Unlock()

	ch := make(chan *DashboardStats, 10)
	ds.subscribers[id] = ch

	// Send current stats immediately
	go func() {
		ch <- ds.GetStats()
	}()

	return ch
}

// Unsubscribe removes a subscriber
func (ds *DashboardService) Unsubscribe(id string) {
	ds.subMu.Lock()
	defer ds.subMu.Unlock()

	if ch, exists := ds.subscribers[id]; exists {
		close(ch)
		delete(ds.subscribers, id)
	}
}

// startRealTimeUpdates runs in background to generate periodic updates
func (ds *DashboardService) startRealTimeUpdates() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	activityTicker := time.NewTicker(5 * time.Second)
	defer activityTicker.Stop()

	for {
		select {
		case <-ticker.C:
			ds.updateStats()
			ds.broadcastStats()

		case <-activityTicker.C:
			ds.addRandomActivity()

		case <-ds.stopChan:
			return
		}
	}
}

// updateStats generates new random stats
func (ds *DashboardService) updateStats() {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	// Simulate gradual changes
	ds.stats.UsersOnline += rand.Intn(6) - 3 // -3 to +3
	if ds.stats.UsersOnline < 10 {
		ds.stats.UsersOnline = 10
	}
	if ds.stats.UsersOnline > 100 {
		ds.stats.UsersOnline = 100
	}

	ds.stats.PageViews += rand.Intn(50) + 10
	ds.stats.ViewsChange = float64(rand.Intn(50)) + 5.0

	// Response time simulation
	responseMs := rand.Intn(200) + 20
	ds.stats.ResponseTime = fmt.Sprintf("%dms", responseMs)
	if responseMs < 50 {
		ds.stats.ResponseStatus = "Excellent"
	} else if responseMs < 100 {
		ds.stats.ResponseStatus = "Good"
	} else {
		ds.stats.ResponseStatus = "Fair"
	}

	// System metrics
	ds.stats.CPUUsage = float64(rand.Intn(60)) + 20
	ds.stats.DiskUsage = float64(rand.Intn(50)) + 25
	ds.stats.NetworkIO = rand.Intn(50) + 10
	ds.stats.NetworkIOPercent = float64(ds.stats.NetworkIO) * 2

	// Activity metrics (simulate varying usage)
	ds.stats.SearchActivity = float64(rand.Intn(50)) + 25
	ds.stats.DashboardActivity = float64(rand.Intn(40)) + 40
	ds.stats.FormsActivity = float64(rand.Intn(40)) + 10
	ds.stats.ComponentsActivity = float64(rand.Intn(30)) + 50

	// Health status
	if ds.stats.CPUUsage > 80 || ds.stats.MemoryPercent > 85 {
		ds.stats.HealthStatus = "warning"
	} else if ds.stats.CPUUsage > 95 || ds.stats.MemoryPercent > 95 {
		ds.stats.HealthStatus = "critical"
	} else {
		ds.stats.HealthStatus = "healthy"
	}

	ds.stats.LastUpdate = time.Now()
}

// addRandomActivity generates random activity items
func (ds *DashboardService) addRandomActivity() {
	activities := []ActivityItem{
		{
			ID:        generateID(),
			Type:      "search",
			Message:   "Nova busca por 'Go framework'",
			Timestamp: time.Now(),
			Icon:      "search",
			Color:     "text-blue-600",
		},
		{
			ID:        generateID(),
			Type:      "user",
			Message:   "Novo usuário conectado",
			Timestamp: time.Now(),
			Icon:      "users",
			Color:     "text-green-600",
		},
		{
			ID:        generateID(),
			Type:      "form",
			Message:   "Formulário de contato enviado",
			Timestamp: time.Now(),
			Icon:      "email",
			Color:     "text-purple-600",
		},
		{
			ID:        generateID(),
			Type:      "system",
			Message:   "Backup automático concluído",
			Timestamp: time.Now(),
			Icon:      "check",
			Color:     "text-green-600",
		},
		{
			ID:        generateID(),
			Type:      "dashboard",
			Message:   "Dashboard visualizado por usuário",
			Timestamp: time.Now(),
			Icon:      "chart",
			Color:     "text-yellow-600",
		},
	}

	activity := activities[rand.Intn(len(activities))]

	ds.activitiesMu.Lock()
	ds.activities = append(ds.activities, activity)

	// Keep only last 50 activities
	if len(ds.activities) > 50 {
		ds.activities = ds.activities[len(ds.activities)-50:]
	}
	ds.activitiesMu.Unlock()
}

// broadcastStats sends current stats to all subscribers
func (ds *DashboardService) broadcastStats() {
	stats := ds.GetStats()

	ds.subMu.RLock()
	defer ds.subMu.RUnlock()

	for id, ch := range ds.subscribers {
		select {
		case ch <- stats:
		default:
			// Channel is full, remove subscriber
			close(ch)
			delete(ds.subscribers, id)
		}
	}
}

// Stop gracefully shuts down the service
func (ds *DashboardService) Stop() {
	close(ds.stopChan)

	ds.subMu.Lock()
	defer ds.subMu.Unlock()

	for _, ch := range ds.subscribers {
		close(ch)
	}
	ds.subscribers = make(map[string]chan *DashboardStats)
}

// Helper functions
func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
