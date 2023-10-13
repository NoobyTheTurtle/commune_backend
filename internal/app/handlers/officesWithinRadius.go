package handlers

import (
	"commune_backend/internal/app/models/sub_models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetOfficesWithinRadius(c *gin.Context) {
	var officesWithDistance sub_models.OfficesWithDistance

	userLat, userLng, radius, err := paramsWithinRadius(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := officesWithDistance.GetWithinRadius(h.DB, userLat, userLng, radius); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &officesWithDistance)
}
