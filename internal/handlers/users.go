package handlers

import (
	"main/internal/handlers/api"
	"main/internal/services"

	"github.com/gin-gonic/gin"
)

func UserListHandler(c *gin.Context) {
	req := services.UserQueryDto{}
	if err := BindAny(c, &req); err != nil {
		api.Error(c, err)
	}
	resp, err := services.UserList(c, req)
	if err != nil {
		api.Error(c, err)
		return
	}
	api.Success(c, gin.H{
		"users": resp.Users,
		"total": resp.Total,
	})
	return
}

func UserCreateHandler(ctx *gin.Context) {

	return
}
