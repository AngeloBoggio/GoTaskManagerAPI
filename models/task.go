package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	Title 	string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed bool `json:"completed"`
}

