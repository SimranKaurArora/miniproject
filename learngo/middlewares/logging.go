package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
)

func LoggingMiddleware(c *gin.Context) {
	log.Println("request recieved", c.Request.Method, c.Request.URL.Path)
	c.Next()
}
