package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
