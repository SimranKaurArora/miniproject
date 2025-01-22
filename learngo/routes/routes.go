package routes

import (
	"github.com/gin-gonic/gin"
	"learngo/middlewares"
)
import "learngo/controllers"

func SetUpRoutes(router *gin.Engine) {

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books", controllers.CreateBook)
	router.GET("books/search", controllers.SearchBook)
	protected := router.Group("/books")
	protected.Use(middlewares.BasicAuthMiddleware)
	protected.PUT("/:id", controllers.UpdateBook)
	protected.DELETE("/:id", controllers.DeleteBook)

}
