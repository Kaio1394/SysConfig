package middleware

import (
	"SysConfig/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		if utils.RequestValidateToken(token) == false {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		c.Next()
	}
}
