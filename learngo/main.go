package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learngo/middlewares"
	"learngo/models"
	"learngo/routes"
	"log"
)

var db *gorm.DB
var err error

func main() {
	dsn := "root:newpassword@tcp(127.0.0.1:3306)/book_management_system?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	models.InitializeDB(db)
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal("Error during AutoMigrate: ", err)
	}
	router := gin.Default()
	router.Use(middlewares.LoggingMiddleware)
	routes.SetUpRoutes(router)
	router.Run(":8080")

}
