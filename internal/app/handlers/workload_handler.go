package handlers

import (
	"commune_backend/internal/app/handlers/params"
	"commune_backend/internal/app/models"
	"commune_backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetWorkloadByOffice godoc
// @Summary Получить загруженность отделения банка по ID
// @Tags Загруженность
// @Accept	json
// @Produce	json
// @Param officeId path uint true "Office ID"
// @Success 200 {object} []models.Workload
// @Failure	400	{object} utils.HttpError
// @Failure	500	{object} utils.HttpError
// @Router /workloads/{officeId} [get]
func (h handler) GetWorkloadByOffice(c *gin.Context) {
	var workloads []models.Workload

	officeId, err := params.ParseParamOfficeId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewHttpError(err))
		return
	}

	if result := h.DB.Order("date_time asc").Find(&workloads, "office_id = ?", officeId); result.Error != nil {
		c.JSON(http.StatusInternalServerError, utils.NewHttpError(result.Error))
	}

	c.JSON(http.StatusOK, &workloads)
}
