package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		fmt.Println("role", role)
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}
		next(c)
	}
}
