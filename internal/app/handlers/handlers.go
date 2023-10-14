package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(routes *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes.GET("/ping", h.Ping)
	routes.GET("/offices", h.GetOfficesWithinRadius)
	routes.GET("/atms", h.GetAtmsWithinRadius)
	routes.POST("/ticket", h.PickTicket)
	routes.GET("/tickets", h.GetTickets)
}
