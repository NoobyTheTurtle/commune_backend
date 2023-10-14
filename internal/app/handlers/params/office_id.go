package params

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseParamsOfficeId(c *gin.Context) (int64, error) {
	officeIdStr := c.Query("officeId")
	userId, err := strconv.ParseInt(officeIdStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, err
}
