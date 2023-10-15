package app

import (
	_ "commune_backend/docs"
	"commune_backend/internal/app/config"
	"commune_backend/internal/app/db"
	"commune_backend/internal/app/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

func Start(c *config.Config) error {
	gin.SetMode(c.GinMode)
	r := gin.Default()
	dbUrl := fmt.Sprintf("postgres://postgres:%s@postgres:%s/%s?sslmode=disable", c.DBPassword, c.DBPort, c.DBName)
	lDb := logger.Default.LogMode(c.DBLogLevel)
	database, err := db.Init(dbUrl, lDb)
	if err != nil {
		return err
	}

	handlers.InitRoutes(r, database)

	if err := r.Run(fmt.Sprintf(":%s", c.ServerPort)); err != nil {
		return err
	}

	return nil
}
