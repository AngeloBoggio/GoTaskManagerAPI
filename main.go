package main

import (
	"github.com/AngeloBoggio/GoTaskManagerAPI/config"
	"github.com/AngeloBoggio/GoTaskManagerAPI/handlers"
	"github.com/AngeloBoggio/GoTaskManagerAPI/models"
	"github.com/gin-gonic/gin"
)

func main() {
    // Set Gin to release mode
    gin.SetMode(gin.ReleaseMode)
    
    // Initialize Gin router
    router := gin.Default()

    // Set trusted proxies
    router.SetTrustedProxies([]string{"127.0.0.1"})

    // Connect to the database
    config.ConnectDatabase()
    
    // Auto-migrate the Task model
    config.DB.AutoMigrate(&models.Task{})

    // Define routes
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    router.GET("/tasks", handlers.GetTasks)
    router.POST("/tasks", handlers.CreateTask)

    // Run the Gin server
    router.Run(":8080")
}
