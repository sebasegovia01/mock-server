package main

import (
	"mock-server/config"
	"mock-server/controllers"
	"mock-server/middleware"
	"mock-server/pkg/logger"
	"mock-server/routes"
	"mock-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Global variables for testing
var loadConfigFunc = config.LoadConfig
var setupServerFunc = setupServer

// EngineRunner interface for testing
type EngineRunner interface {
	Run(addr ...string) error
}

type engineRunnerAdapter struct {
	*gin.Engine
}

func adaptGinToEngineRunner(engine *gin.Engine) EngineRunner {
	return engineRunnerAdapter{engine}
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	if err := run(loadConfigFunc, setupServerFunc); err != nil {
		logger.ErrorLogger.Fatalf("%v", err)
	}
}

func run(
	loadConfigFunc func() (*config.Config, error),
	setupServerFunc func(*config.Config) (EngineRunner, error),
) error {
	cfg, err := loadConfigFunc()
	if err != nil {
		return err
	}

	if err := initializeData(cfg); err != nil {
		return err
	}

	r, err := setupServerFunc(cfg)
	if err != nil {
		return err
	}

	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	logger.InfoLogger.Printf("Server running on port: %s", port)
	return r.Run(":" + port)
}

func setupServer(cfg *config.Config) (EngineRunner, error) {
	router := gin.New() // Use New() instead of Default()

	// Basic middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ResponseWrapperMiddleware())

	// Handler for not found routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	})

	// Configure routes
	routes.SetupRoutes(router, cfg)

	return adaptGinToEngineRunner(router), nil
}

func initializeData(cfg *config.Config) error {
	usersData, err := services.GetS3Data(cfg)
	if err != nil {
		return err
	}

	controllers.InitializeData(usersData)
	return nil
}
