package handlers

import (
	"commune_backend/internal/app/handlers/params"
	"commune_backend/internal/app/models"
	"commune_backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// GetOffice godoc
// @Summary Получить отделение банка по ID
// @Tags Отделения банков
// @Accept	json
// @Produce	json
// @Param id path uint true "Office ID"
// @Success 200 {object} models.Office
// @Failure	404	{object} utils.HttpError
// @Failure	400	{object} utils.HttpError
// @Failure	500	{object} utils.HttpError
// @Router /offices/{id} [get]
func (h handler) GetOffice(c *gin.Context) {
	var office models.Office

	id, err := params.ParseParamId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewHttpError(err))
		return
	}

	if result := h.DB.First(&office, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.NewHttpError(errors.New("office not found")))
		} else {
			c.JSON(http.StatusInternalServerError, utils.NewHttpError(result.Error))
		}
		return
	}

	c.JSON(http.StatusOK, &office)
}
