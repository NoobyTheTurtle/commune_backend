package handlers

import (
	"commune_backend/internal/app/filters"
	"commune_backend/internal/app/handlers/params"
	"commune_backend/internal/app/models/support_models"
	"commune_backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAtmsWithinRadius godoc
// @Summary Получить банкоматы
// @Description Получить банкоматы в радиусе пользователя
// @Tags Банкоматы
// @Accept			json
// @Produce	json
// @Param userLat query float64 true "User latitude"
// @Param userLng query float64 true "User longitude"
// @Param radius query float64 true "Radius in meters"
// @Param isImmobile query bool false "Filter by immobile"
// @Param isWithdrawal query bool false "Filter by withdrawal"
// @Param isReplenishment query bool false "Filter by replenishment"
// @Success 200 {object} support_models.AtmsWithRadius
// @Failure	400	{object} utils.HttpError
// @Failure	500	{object} utils.HttpError
// @Router /atms [get]
func (h handler) GetAtmsWithinRadius(c *gin.Context) {
	var atmsWithDistance support_models.AtmsWithRadius
	var ga params.GeoArea

	err := ga.ParseParamsGeoArea(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewHttpError(err))
		return
	}

	f := filters.Filters{
		filters.FilterAtmsByImmobile(c.Query("isImmobile")),
		filters.FilterAtmsByWithdrawal(c.Query("isWithdrawal")),
		filters.FilterAtmsByReplenishment(c.Query("isReplenishment")),
	}

	if err := atmsWithDistance.GetWithinRadius(h.DB, &ga, &f); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewHttpError(err))
		return
	}

	c.JSON(http.StatusOK, &atmsWithDistance)
}
