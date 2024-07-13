package handlers

import (
	"errors"
	"net/http"

	"github.com/AngeloBoggio/GoTaskManagerAPI/config"
	"github.com/AngeloBoggio/GoTaskManagerAPI/middleware"
	"github.com/AngeloBoggio/GoTaskManagerAPI/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GetTasks handles GET requests to retrieve tasks for the authenticated user
func GetTasks(c *gin.Context) {
    // Retrieve user ID from context (assuming it's set by middleware)
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var tasks []models.Task
    // Query tasks associated with the authenticated user
    if err := config.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}


func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Retrieve user ID from context (assuming it's set by middleware)
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Assign user ID to the task
    task.UserID = userID.(uint) // Assuming task struct has a UserID field

    // Create the task in the database
    if err := config.DB.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return created task
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
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}


func SignUp(c *gin.Context) {
    var user models.User
    // Bind JSON to user model first
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
        return
    }

    // Check if the user already exists
    var existingUser models.User
    // Find the first user with name of user
    if err := config.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
        // If no error, user exists
        c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
        return
    } else if !errors.Is(err, gorm.ErrRecordNotFound) {
        // If the error is not ErrRecordNotFound, then it's an unexpected error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    user.Password = string(hashedPassword)
    
    // Create new user
    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    // Successfully created user
    c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context){
    
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    var existingUser models.User
    if err := config.DB.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }


    // Compare the hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }


     // Generate token if login is successful
     token, err := middleware.GenerateToken(existingUser.ID)
     if err != nil {
         c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
         return 
     }
 
     c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user":  existingUser, // This includes the tasks associated with the user
    })
}
