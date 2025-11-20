package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ramith/service-now-mapping-service/internal/config"
	"github.com/ramith/service-now-mapping-service/internal/handlers"
	"github.com/ramith/service-now-mapping-service/internal/models"
	"github.com/ramith/service-now-mapping-service/internal/repository"
	"github.com/ramith/service-now-mapping-service/internal/service"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := initDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate database schema
	if err := db.AutoMigrate(&models.APIMapping{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize layers
	repo := repository.NewAPIMappingRepository(db)
	svc := service.NewAPIMappingService(repo)
	handler := handlers.NewAPIMappingHandler(svc)

	// Set up router
	router := setupRouter(handler)

	// Create HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on %s:%s", cfg.Server.Host, cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited successfully")
}

func initDatabase(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Get underlying SQL database for connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func setupRouter(handler *handlers.APIMappingHandler) *gin.Engine {
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", handler.HealthCheck)

	// Mapping routes
	mappings := router.Group("/mappings")
	{
		mappings.POST("", handler.CreateMapping)
		mappings.GET("", handler.GetAllMappings)
		mappings.GET("/sys-id/:sys_id", handler.GetMappingBySysID)
		mappings.GET("/api-id/:api_id", handler.GetMappingByAPIID)
		mappings.GET("/search", handler.GetMappingByNameAndVersion)
		mappings.PUT("/:id", handler.UpdateMapping)
		mappings.DELETE("/:id", handler.DeleteMapping)
	}

	return router
}
