package models // This package is called models and it is used to store the data models for the application

type Task struct {
	ID 		uint   `json:"id" gorm:"primary_key"`  // This is an unsigned integer that represents the ID of the task but whats this syntax?
	Title 	string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}
