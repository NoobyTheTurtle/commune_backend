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
	routes.GET("/bank_branches", h.GetBankBranches)
}
