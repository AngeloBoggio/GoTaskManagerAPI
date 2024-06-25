package models // This package is called models and it is used to store the data models for the application

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title 	string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed bool `json:"completed"`
}

