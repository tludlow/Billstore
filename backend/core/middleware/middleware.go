package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test1/core/jwt"
)

func Auth(c *gin.Context) {
	accessToken, err := c.Request.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access_token"})
		c.Abort()
		return
	}
	tokenId, err := jwt.Validate(accessToken.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("userId", tokenId.String())
	c.Next()
}