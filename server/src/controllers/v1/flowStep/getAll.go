package flowStep

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"server/src/utils"
	"strconv"
)

// GetAllFlowStep godoc
// @Security Session Token
// @Summary Retrieve all flow step
// @Tags flow steps
// @Param page query int false "Page" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Produce json
// @Success 200 {object} dto.PaginatedResponse{data=dto.FlowStep} "OK"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/flowsteps [get]
func GetAllFlowStep(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	var flowSteps []models.FlowStep

	result := config.DB.Find(&flowSteps)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed find flows steps."})
	}

	var DTOflowsStep []dto.FlowStep

	for _, flowStep := range flowSteps {
		DTOflowsStep = append(DTOflowsStep, flowStep.GetSimpleStep())
	}

	c.JSON(http.StatusOK, utils.Paginate(DTOflowsStep, page, pageSize, len(DTOflowsStep)))
}
