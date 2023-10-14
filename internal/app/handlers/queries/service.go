package queries

import (
	"commune_backend/internal/app/models"
	"commune_backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func ParseQueryService(c *gin.Context) (string, error) {
	service := c.Query("service")
	condition := utils.Contains(models.ServicesList, service)
	if condition {
		return service, nil
	}
	return "", errors.New("service invalid")
}
