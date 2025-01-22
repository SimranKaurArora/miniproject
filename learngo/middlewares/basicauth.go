package middlewares

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func validateCredentials(username, password string) bool {
	expectedUsername := "admin"
	expectedPassword := "password"
	return username == expectedUsername && password == expectedPassword

}
func BasicAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Empty authorization",
		})
		c.Abort()
		return
	}
	if !strings.HasPrefix(authHeader, "Basic ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Invalid Authorization header",
		})
		c.Abort()
		return
	}

	credentials := strings.TrimPrefix(authHeader, "Basic ")
	decodedCredentials, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Invalid base64 encoding in Authorization header",
		})
		c.Abort()
		return
	}
	parts := strings.SplitN(string(decodedCredentials), ":", 2)
	if len(parts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Invalid credentials format",
		})
		c.Abort()
		return
	}
	username, password := parts[0], parts[1]

	if !validateCredentials(username, password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		c.Abort()
		return
	}
	c.Next()

}
