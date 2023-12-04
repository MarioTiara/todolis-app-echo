package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPageAndLimit(c *gin.Context) (int, int) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	return page, limit
}
