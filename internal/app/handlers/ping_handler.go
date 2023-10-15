package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping godoc
// @Summary Ping API
// @Description Ping API for checking server status
// @Tags Check
// @ID ping
// @Produce	json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (h handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
