package main

import (
	"log"
	"my-procurement-system/internal/config"
	"my-procurement-system/internal/routes"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Initialize database
	db := config.InitDB()

	// Run database migration
	if err := config.Migrate(db); err != nil {
		log.Fatal("❌ Database migration failed")
	}

	// Setup Gin
	r := gin.Default()

	// setup cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"*", "Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	// Register routes
	api := r.Group("/api")
	routes.RegisterItemRoutes(api, db)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}
	log.Printf("🚀 Server running on port %s", port)
	r.Run(port)
}
