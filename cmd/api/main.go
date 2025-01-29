// cmd/api/main.go
package main

import (
	"log"

	"server/internal/config"
	"server/internal/database"
	"server/internal/handlers"
	"server/internal/repositories"
	"server/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Run database migrations
	runMigrations(db)

	// Initialize repositories
	categoryRepo := repositories.NewCategoryRepo(db)
	serviceRepo := repositories.NewServiceRepo(db)

	// Initialize services
	categoryService := services.NewCategoryService(categoryRepo)
	serviceService := services.NewServiceService(serviceRepo, categoryRepo)

	// Initialize handlers
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	serviceHandler := handlers.NewServiceHandler(serviceService)

	// Create router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		err := db.Ping()
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Database connection failed",
			})
			return
		}
		c.JSON(200, gin.H{
			"status":   "ok",
			"database": "connected",
		})
	})

	// Category routes
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories", categoryHandler.ListCategories)
	
	// IMPORTANT: The more specific route comes first
	r.GET("/categories/:id/services", serviceHandler.ListServicesByCategory)
	
	// General category routes
	r.GET("/categories/:id", categoryHandler.GetCategory)
	r.PUT("/categories/:id", categoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	// Service routes
	r.POST("/services", serviceHandler.CreateService)
	r.GET("/services", serviceHandler.ListServices)
	r.GET("/services/:id", serviceHandler.GetService)
	r.PUT("/services/:id", serviceHandler.UpdateService)
	r.DELETE("/services/:id", serviceHandler.DeleteService)

	// Start server
	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}

func runMigrations(db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatal("Migration setup failed:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", 
		driver,
	)
	if err != nil {
		log.Fatal("Migration initialization failed:", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}

	version, dirty, err := m.Version()
	if err == nil {
		log.Printf("Database migrated to version: %d (dirty: %v)", version, dirty)
	}
}