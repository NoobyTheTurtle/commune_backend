package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func InitRoutes(routes *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes.GET("/offices", h.GetOfficesWithinRadius)
	routes.GET("/offices/:id", h.GetOffice)

	routes.GET("/atms", h.GetAtmsWithinRadius)
	routes.GET("/atms/:id", h.GetAtm)

	routes.GET("/tickets", h.GetTickets)
	routes.POST("/ticket", h.PickTicket)

	routes.GET("/workloads/:officeId", h.GetWorkloadByOffice)

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
