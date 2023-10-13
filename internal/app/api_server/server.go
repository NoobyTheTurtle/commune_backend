package api_server

import (
	"commune_backend/internal/app/config"
	"commune_backend/internal/app/db"
	"commune_backend/internal/app/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start(c *config.Config) error {
	gin.SetMode(c.GinMode)
	r := gin.Default()
	dbUrl := fmt.Sprintf("postgres://postgres:%s@postgres:%s/%s?sslmode=disable", c.DBPassword, c.DBPort, c.DBName)
	database, err := db.Init(dbUrl)
	if err != nil {
		return err
	}

	handlers.RegisterRoutes(r, database, c)

	if err := r.Run(fmt.Sprintf(":%s", c.ServerPort)); err != nil {
		return err
	}

	return nil
}
