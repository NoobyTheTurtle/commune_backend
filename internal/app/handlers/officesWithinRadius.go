package handlers

import (
	"commune_backend/internal/app/filters"
	"commune_backend/internal/app/handlers/params"
	"commune_backend/internal/app/models/support_models"
	"commune_backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOfficesWithinRadius godoc
// @Summary Получить отделения банков
// @Description Получить отделения банков в радиусе пользователя
// @Tags Отделения банков
// @Accept	json
// @Produce	json
// @Param userLat query float64 true "User latitude"
// @Param userLng query float64 true "User longitude"
// @Param radius query float64 true "Radius in meters"
// @Param isImmobile query bool false "Filter by immobile"
// @Param isLegalPerson query bool false "Filter by legal person"
// @Param isIndividualPerson query bool false "Filter by individual person"
// @Param isPrime query bool false "Filter by prime"
// @Param isOpen query bool false "Filter by open"
// @Param isWithdrawal query bool false "Filter by withdrawal"
// @Param isReplenishment query bool false "Filter by replenishment"
// @Success 200 {object} support_models.OfficesWithRadius
// @Failure	400	{object} utils.HttpError
// @Failure	500	{object} utils.HttpError
// @Router /offices [get]
func (h handler) GetOfficesWithinRadius(c *gin.Context) {
	var officesWithDistance support_models.OfficesWithRadius
	var ga params.GeoArea

	err := ga.ParseParamsGeoArea(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewHttpError(err))
		return
	}

	f := filters.Filters{
		filters.FilterOfficesByLegalPerson(c.Query("isLegalPerson")),
		filters.FilterOfficesByIndividualPerson(c.Query("isIndividualPerson")),
		filters.FilterOfficesByImmobile(c.Query("isImmobile")),
		filters.FilterOfficesByPrime(c.Query("isPrime")),
		filters.FilterOfficesByOpen(c.Query("isOpen"), c.Query("isLegalPerson")),
		filters.FilterOfficesByWithdrawal(c.Query("isWithdrawal")),
		filters.FilterOfficesByReplenishment(c.Query("isReplenishment")),
	}

	if err := officesWithDistance.GetWithinRadius(h.DB, &ga, &f); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewHttpError(err))
		return
	}

	c.JSON(http.StatusOK, &officesWithDistance)
}
