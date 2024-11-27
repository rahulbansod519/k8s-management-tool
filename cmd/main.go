package main

import (
	"k8s-management-tool/internal/handlers"
	"k8s-management-tool/internal/middleware"
	"log"

	"k8s-management-tool/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatalf("Error Loading environment variables: %v", err)
	}
	config.ConnectDatabase()
	router := gin.Default()

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", handlers.Register)
		authRoutes.POST("/login", handlers.Login)
	}

	protectedRoutes := router.Group("/cluster")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.POST("/register", handlers.RegisterCluster)
		protectedRoutes.GET("/status", handlers.ClusterStatus)
	}

	log.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file-found")
	}
	return err
}
