package helpers

import (
	"github.com/gin-gonic/gin"
)

func BindAndValidate(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		ResponseJSON(c, 400, "Validation Failed", ValidationErrors(err))
		return false
	}
	return true
}