package handlers

import (
	"commune_backend/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h handler) GetTickets(c *gin.Context) {
	var tickets []models.Ticket

	userId, err := parseParamsUserId(c)
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
	userId, err := parseParamsUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	officeId, err = parseParamsOfficeId(c)
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

func parseParamsUserId(c *gin.Context) (int64, error) {
	userIdStr := c.Query("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, err
}

func parseParamsOfficeId(c *gin.Context) (int64, error) {
	officeIdStr := c.Query("officeId")
	userId, err := strconv.ParseInt(officeIdStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, err
}
