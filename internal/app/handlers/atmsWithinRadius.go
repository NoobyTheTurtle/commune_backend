package handlers

import (
	"commune_backend/internal/app/models/sub_models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetAtmsWithinRadius(c *gin.Context) {
	var atmsWithDistance sub_models.AtmsWithDistance

	userLat, userLng, radius, err := paramsWithinRadius(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := atmsWithDistance.GetWithinRadius(h.DB, userLat, userLng, radius); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &atmsWithDistance)
}
