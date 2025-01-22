package controllers

import (
	"github.com/gin-gonic/gin"
	"learngo/models"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	if err := models.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}
func GetBook(c *gin.Context) {
	id := c.Param("id")
	bookid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "invalid id",
		})
		return
	}
	var book models.Book
	if err := models.DB.First(&book, bookid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "Book not found",
		})
		return

	}
	c.JSON(http.StatusOK, book)

}
func CreateBook(c *gin.Context) {
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid request body",
		})
		return
	}
	err = models.DB.Create(&book).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "Cannot create new book",
		})
		return
	}
	c.JSON(http.StatusOK, book)
}
func UpdateBook(c *gin.Context) {

	id := c.Param("id")
	bookid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid id ",
		})
		return
	}
	var book models.Book
	err = models.DB.First(&book, bookid).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Book doesn't exists",
		})
		return
	}
	err = c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "Invalid request data ",
		})
		return

	}
	if err := models.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Cannot update the book",
		})
		return
	}
	c.JSON(http.StatusOK, book)
}
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid id",
		})
		return
	}
	var book models.Book
	err = models.DB.First(&book, bookid).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Book doesn't exists",
		})
		return
	}
	if err := models.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Error deleting book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
func SearchBook(c *gin.Context) {
	var books []models.Book
	title := c.DefaultQuery("title", "")
	author := c.DefaultQuery("author", "")
	year := c.DefaultQuery("year", "")
	query := models.DB.Model(&models.Book{})
	if title != "" {
		query = query.Where("title like ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("author like ?", "%"+author+"%")
	}
	if year != "" {
		query.Where("publication_year like ?", "%"+year+"%")
	}
	if err := query.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Error retrieving books"})
		return
	}
	c.JSON(http.StatusOK, books)
}
