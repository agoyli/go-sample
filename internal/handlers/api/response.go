package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data gin.H) {
	data["success"] = true
	c.JSON(http.StatusOK, data)
}

func Error(c *gin.Context, err error) {
	data := gin.H{}
	data["success"] = false
	data["errors"] = errToResponse(err)

	c.JSON(http.StatusBadRequest, data)
}

type errorResponse struct {
	Code  string `json:"code"`
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func errToResponse(err error) []errorResponse {
	return []errorResponse{
		{
			Msg: err.Error(),
		},
	}
}
