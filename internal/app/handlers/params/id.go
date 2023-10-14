package params

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseParamId(c *gin.Context) (uint, error) {
	userIdStr := c.Param("id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(userId), err
}
