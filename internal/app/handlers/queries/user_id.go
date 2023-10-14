package queries

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseQueryUserId(c *gin.Context) (int64, error) {
	userIdStr := c.Query("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, err
}
