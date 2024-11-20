package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"server/src/utils"
	"strconv"
)

// GetAllActions godoc
// @Security Session Token
// @Summary Retrieve all actions
// @Tags actions
// @Param page query int false "Page" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Produce json
// @Success 200 {object} dto.PaginatedResponse{data=dto.Action} "OK"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/action [get]
func GetAllActions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	var actions []models.Action

	result := config.DB.Find(&actions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed find actions."})
	}

	var DTOactions []dto.Action

	for _, action := range actions {
		DTOactions = append(DTOactions, action.GetSimpleAction())
	}

	c.JSON(http.StatusOK, utils.Paginate(DTOactions, page, pageSize, len(DTOactions)))
}
