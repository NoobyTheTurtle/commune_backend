package handlers

import (
	"commune_backend/internal/app/handlers/params"
	"commune_backend/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetTickets(c *gin.Context) {
	var tickets []models.Ticket

	userId, err := params.ParseParamsUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.DB.Where("user_id = ?", userId).Find(&tickets).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &tickets)
}

func (h handler) PickTicket(c *gin.Context) {
	var officeId, userId int64
	userId, err := params.ParseParamsUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	officeId, err = params.ParseParamsOfficeId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket := models.Ticket{OfficeId: officeId, UserId: userId}

	err = h.DB.Create(&ticket).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, &ticket)
}
