package handlers

import (
	"commune_backend/internal/app/filters"
	"commune_backend/internal/app/handlers/params"
	"commune_backend/internal/app/models/support_models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetOfficesWithinRadius(c *gin.Context) {
	var officesWithDistance support_models.OfficesWithRadius
	var ga params.GeoArea

	err := ga.ParseParamsGeoArea(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f := filters.Filters{
		filters.FilterOfficesByLegalPerson(c.Query("isLegalPerson")),
		filters.FilterOfficesByIndividualPerson(c.Query("isIndividualPerson")),
		filters.FilterOfficesByImmobile(c.Query("isImmobile")),
		filters.FilterOfficesByPrime(c.Query("isPrime")),
		filters.FilterOfficesByOpen(c.Query("isOpen"), c.Query("isLegalPerson")),
	}

	if err := officesWithDistance.GetWithinRadius(h.DB, &ga, &f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &officesWithDistance)
}
