package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func paramsWithinRadius(c *gin.Context) (userLat float64, userLng float64, radius float64, err error) {
	userLatStr := c.Query("userLat")
	userLngStr := c.Query("userLng")
	radiusStr := c.Query("radius")

	userLat, err = strconv.ParseFloat(userLatStr, 64)
	if err != nil {
		return
	}

	userLng, err = strconv.ParseFloat(userLngStr, 64)
	if err != nil {
		return
	}

	radius, err = strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		return
	}

	return
}
