package api_server

import (
	"commune_backend/internal/app/config"
	"commune_backend/internal/app/db"
	"commune_backend/internal/app/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start(config *config.Config) error {
	gin.SetMode(config.GinMode)
	r := gin.Default()
	dbUrl := fmt.Sprintf("postgres://postgres:%s@postgres:%s/%s?sslmode=disable", config.DBPassword, config.DBPort, config.DBName)
	database, err := db.Init(dbUrl)
	if err != nil {
		return err
	}

	handlers.RegisterRoutes(r, database)

	if err := r.Run(fmt.Sprintf(":%s", config.ServerPort)); err != nil {
		return err
	}

	return nil
}
