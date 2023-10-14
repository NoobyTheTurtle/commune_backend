package queries

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type GeoArea struct {
	UserLat float64
	UserLng float64
	Radius  float64
}

func (ga *GeoArea) ParseQueryGeoArea(c *gin.Context) (err error) {
	ga.UserLat, err = strconv.ParseFloat(c.Query("userLat"), 64)
	if err != nil {
		return
	}

	ga.UserLng, err = strconv.ParseFloat(c.Query("userLng"), 64)
	if err != nil {
		return
	}

	ga.Radius, err = strconv.ParseFloat(c.Query("radius"), 64)
	if err != nil {
		return
	}

	return
}
