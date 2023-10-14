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

// GetAtm godoc
// @Summary Получить банкомат по ID
// @Tags Банкоматы
// @Accept	json
// @Produce	json
// @Param id path uint true "Atm ID"
// @Success 200 {object} models.Atm
// @Failure	404	{object} utils.HttpError
// @Failure	400	{object} utils.HttpError
// @Failure	500	{object} utils.HttpError
// @Router /atms/{id} [get]
func (h handler) GetAtm(c *gin.Context) {
	var atm models.Atm

	id, err := params.ParseParamId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewHttpError(err))
		return
	}

	if err := atm.Get(h.DB, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.NewHttpError(errors.New("atm not found")))
		} else {
			c.JSON(http.StatusInternalServerError, utils.NewHttpError(err))
		}
		return
	}

	c.JSON(http.StatusOK, &atm)
}
