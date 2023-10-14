package handlers

import (
	"commune_backend/internal/app/handlers/queries"
	"commune_backend/internal/app/models"
	"commune_backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTickets godoc
// @Summary Получить талоны пользователя
// @Tags Талоны
// @Accept	json
// @Produce	json
// @Param userId query uint true "User ID"
// @Success 200 {object} []models.Ticket
// @Failure	400	{object} utils.HttpError
// @Failure	500	{object} utils.HttpError
// @Router /tickets [get]
func (h handler) GetTickets(c *gin.Context) {
	var tickets []models.Ticket

	userId, err := queries.ParseQueryUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewHttpError(err))
		return
	}

	err = h.DB.Where("user_id = ?", userId).
		Order("created_at desc").
		Limit(20).
		Find(&tickets).
		Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewHttpError(err))
		return
	}
	c.JSON(http.StatusOK, &tickets)
}

// PickTicket godoc
// @Summary Получить талон для отделения
// @Tags Талоны
// @Accept	json
// @Produce	json
// @Param userId query uint true "User ID"
// @Param officeId query uint true "Office ID"
// @Param service query string true "Filter by services" Enums(кредит, карта, ипотека, автокредит, вклад и счет, платежи и переводы, страхование, другие услуги)
// @Success 201 {object} models.Ticket
// @Failure	400	{object} utils.HttpError
// @Failure	500	{object} utils.HttpError
// @Router /ticket [post]
func (h handler) PickTicket(c *gin.Context) {
	var (
		officeId, userId int64
		service          string
	)
	userId, err := queries.ParseQueryUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewHttpError(err))
		return
	}

	officeId, err = queries.ParseQueryOfficeId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewHttpError(err))
		return
	}

	service, err = queries.ParseQueryService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewHttpError(err))
		return
	}
	ticket := models.NewTicket(officeId, userId, service)

	err = h.DB.Create(&ticket).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewHttpError(err))
	}
	c.JSON(http.StatusCreated, &ticket)
}
