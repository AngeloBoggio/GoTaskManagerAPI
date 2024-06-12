package handlers

import (
	"net/http"

	"github.com/AngeloBoggio/GoTaskManagerAPI/config"
	"github.com/AngeloBoggio/GoTaskManagerAPI/models"
	"github.com/gin-gonic/gin"
)

// GetTasks handles GET requests to retrieve all tasks
func GetTasks(c *gin.Context) {
    var tasks []models.Task
    if err := config.DB.Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

// CreateTask handles POST requests to create a new task
func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := config.DB.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context){
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&models.Task{}).Where("id = ?", c.Param("id")).Updates(task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context){
    var task models.Task
    if err := config.DB.Where("id = ?", c.Params.ByName("id")).Delete(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
}
