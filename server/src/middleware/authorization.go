package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"strconv"
	"strings"
)

func GetAuthozirationInfo(c *gin.Context) (string, int64) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		c.Abort()
		return "", 0
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	userID, err := config.Redis.Get(context.Background(), token).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token or session expired"})
		c.Abort()
		return "", 0
	}

	userId, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token or session expired"})
		c.Abort()
		return "", 0
	}

	var user = models.User{
		ID: int(userId),
	}

	config.DB.Where(&user).First(&user)
	c.Set("user", user)
	c.Set("token", token)
	return token, userId
}

func UUIDAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		GetAuthozirationInfo(c)
		if !c.IsAborted() {
			c.Next()
		}
	}
}
