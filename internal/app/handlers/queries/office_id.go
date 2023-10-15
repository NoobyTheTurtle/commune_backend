package queries

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseQueryOfficeId(c *gin.Context) (int64, error) {
	officeIdStr := c.Query("officeId")
	officeId, err := strconv.ParseInt(officeIdStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return officeId, err
}
