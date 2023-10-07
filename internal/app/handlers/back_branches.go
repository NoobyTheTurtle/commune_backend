package handlers

import (
	"commune_backend/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetBankBranches(c *gin.Context) {
	var bankBranches []models.BankBranch

	if result := h.DB.Find(&bankBranches); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &bankBranches)
}
