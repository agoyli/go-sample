package handlers

import (
	"main/internal/utils"

	"github.com/gin-gonic/gin"
)

func BindAny(c *gin.Context, r interface{}) error {
	if err := c.ShouldBind(r); err != nil {
		return utils.ErrInvalid.SetKey("json").SetComment(err.Error())
	}
	return nil
}
