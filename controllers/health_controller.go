package controllers

import (
	"mock-server/config"
	"mock-server/services"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	startTime time.Time
	config    *config.Config
}

type HealthStatus struct {
	Status    string            `json:"status"`
	Message   string            `json:"message"`
	Details   HealthDetails     `json:"details"`
	Resources ResourcesStatus   `json:"resources"`
	System    SystemInformation `json:"system"`
}

type HealthDetails struct {
	Uptime      string `json:"uptime"`
	Environment string `json:"environment"`
}

type ResourcesStatus struct {
	S3Storage ComponentStatus `json:"s3_storage"`
}

type ComponentStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SystemInformation struct {
	GoVersion       string `json:"go_version"`
	NumGoroutines   int    `json:"num_goroutines"`
	MemoryAllocated uint64 `json:"memory_allocated_mb"`
}

func NewHealthController(cfg *config.Config) *HealthController {
	return &HealthController{
		startTime: time.Now(),
		config:    cfg,
	}
}

func (c *HealthController) HealthCheck(ctx *gin.Context) {
	health := HealthStatus{
		Status:  "UP",
		Message: "API is healthy",
		Details: HealthDetails{
			Uptime:      time.Since(c.startTime).String(),
			Environment: c.config.Env,
		},
		Resources: c.checkResources(),
		System:    c.getSystemInfo(),
	}

	// Determine overall status
	if health.Resources.S3Storage.Status == "DOWN" {
		health.Status = "DEGRADED"
		health.Message = "Some services are not available"
	}

	ctx.JSON(http.StatusOK, health)
}

func (c *HealthController) checkResources() ResourcesStatus {
	resources := ResourcesStatus{}

	// Check S3 connectivity
	_, err := services.GetS3Data(c.config)
	if err != nil {
		resources.S3Storage = ComponentStatus{
			Status:  "DOWN",
			Message: "Cannot connect to S3: " + err.Error(),
		}
	} else {
		resources.S3Storage = ComponentStatus{
			Status:  "UP",
			Message: "S3 connection successful",
		}
	}

	// Here you could add more resource checks like database
	// resources.Database = ComponentStatus{
	// 	Status:  "UP",
	// 	Message: "Mock database connection successful",
	// }

	return resources
}

func (c *HealthController) getSystemInfo() SystemInformation {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return SystemInformation{
		GoVersion:       runtime.Version(),
		NumGoroutines:   runtime.NumGoroutine(),
		MemoryAllocated: mem.Alloc / (1024 * 1024), // Convert to MB
	}
}
