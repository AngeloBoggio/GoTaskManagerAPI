package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Tasks []Task `json:"tasks"`
}
