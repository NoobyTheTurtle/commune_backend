package handlers

import (
	"commune_backend/internal/app/filters"
	"commune_backend/internal/app/handlers/params"
	"commune_backend/internal/app/models/support_models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetAtmsWithinRadius(c *gin.Context) {
	var atmsWithDistance support_models.AtmsWithRadius
	var ga params.GeoArea

	err := ga.ParseParamsGeoArea(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f := filters.Filters{
		filters.FilterAtmsByImmobile(c.Query("isImmobile")),
	}

	if err := atmsWithDistance.GetWithinRadius(h.DB, &ga, &f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &atmsWithDistance)
}
