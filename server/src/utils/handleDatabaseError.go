package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"server/src/models/dto"
)

func HandleDatabaseError(result *gorm.DB, c *gin.Context) bool {
	if result.Error == nil {
		return true
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "content not found"})
		return false
	}

	c.JSON(http.StatusInternalServerError, dto.RequestError{Error: result.Error.Error()})
	return false
}
