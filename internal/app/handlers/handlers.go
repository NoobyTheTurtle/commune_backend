package handlers

import (
	"commune_backend/internal/app/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB     *gorm.DB
	Config *config.Config
}

func RegisterRoutes(routes *gin.Engine, db *gorm.DB, c *config.Config) {
	h := &handler{
		DB:     db,
		Config: c,
	}
	routes.GET("/ping", h.Ping)
	routes.GET("/offices", h.GetOfficesWithinRadius)
	routes.GET("/atms", h.GetAtmsWithinRadius)
	routes.POST("/ticket", h.PickTicket)
	routes.GET("/tickets", h.GetTickets)
}
