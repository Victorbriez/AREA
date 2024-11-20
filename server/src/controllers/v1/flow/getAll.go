package flow

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"server/src/utils"
	"strconv"
)

// GetAllFlow godoc
// @Security Session Token
// @Summary Retrieve all flows
// @Tags flows
// @Param page query int false "Page" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Produce json
// @Success 200 {object} dto.PaginatedResponse{data=dto.SimpleFlow} "OK"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/flow [get]
func GetAllFlow(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	var flows []models.Flow

	result := config.DB.Find(&flows)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed find flows."})
	}

	var DTOflows []dto.SimpleFlow

	for _, flow := range flows {
		DTOflows = append(DTOflows, flow.GetSimpleFlow())
	}

	c.JSON(http.StatusOK, utils.Paginate(DTOflows, page, pageSize, len(DTOflows)))
}
