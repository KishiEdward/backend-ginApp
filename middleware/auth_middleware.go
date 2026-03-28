package middleware 

import ( 
	"net/http" 
	"os" 
	"strings" 

	"github.com/gin-gonic/gin" 
	"github.com/golang-jwt/jwt/v5" 
) 

func AuthMiddleware() gin.HandlerFunc { 
	return func(c *gin.Context) { 
		authHeader := c.GetHeader("Authorization") 
		if authHeader == "" { 
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ 
				"success":    false, 
				"message":    "Authorization header tidak ditemukan", 
				"error_code": "MISSING_TOKEN", 
			}) 
			return 
		} 