package main

import (
	"github.com/AngeloBoggio/GoTaskManagerAPI/config"
	"github.com/AngeloBoggio/GoTaskManagerAPI/handlers"
	"github.com/AngeloBoggio/GoTaskManagerAPI/middleware"
	"github.com/AngeloBoggio/GoTaskManagerAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
    godotenv.Load()
}

func main() {
    // Set Gin to release mode
    gin.SetMode(gin.ReleaseMode)
    
    // Initialize Gin router
    router := gin.Default()

    // Set trusted proxies
    router.SetTrustedProxies([]string{"127.0.0.1"})

    // Connect to the database
    config.ConnectDatabase()
    
    // Auto-migrate models
    config.DB.AutoMigrate(&models.Task{})
    config.DB.AutoMigrate(&models.User{})
    
     // Public routes
     router.POST("/login", handlers.Login)
     router.POST("/signup", handlers.SignUp)

     // Protected routes
     authorized := router.Group("/")
     authorized.Use(middleware.AuthMiddleware())
     {
         authorized.GET("/tasks", handlers.GetTasks)
         authorized.POST("/tasks", handlers.CreateTask)
         authorized.PUT("/tasks/:id", handlers.UpdateTask)
         authorized.DELETE("/tasks/:id", handlers.DeleteTask)
     }

    // Run the Gin server
    router.Run(":8080")
}
