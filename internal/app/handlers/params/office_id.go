package params

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseParamOfficeId(c *gin.Context) (uint, error) {
	paramStr := c.Param("officeId")
	value, err := strconv.ParseUint(paramStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(value), err
}
