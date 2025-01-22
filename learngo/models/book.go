package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publication_year"`
}

var DB *gorm.DB

func InitializeDB(db *gorm.DB) {
	DB = db
}
