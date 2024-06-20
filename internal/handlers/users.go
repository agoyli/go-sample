package handlers

import (
	"main/internal/handlers/api"
	"main/internal/services"
	"main/internal/services/validator"

	"github.com/gin-gonic/gin"
)

func UserListHandler(c *gin.Context) {
	// bind request
	req := services.UserQueryDto{}
	if err := BindAny(c, &req); err != nil {
		api.Error(c, err)
		return
	}
	// validate
	if err := validator.ValidateUserQuery(req); err != nil {
		api.Error(c, err)
		return
	}

	// handle
	resp, err := services.UserList(c, req)
	if err != nil {
		api.Error(c, err)
		return
	}
	api.Success(c, gin.H{
		"users": resp.Users,
		"total": resp.Total,
	})
}

func UserCreateHandler(c *gin.Context) {
	// bind request
	req := services.UserCreateDto{}
	if err := BindAny(c, &req); err != nil {
		api.Error(c, err)
		return
	}
	// validate
	if err := validator.ValidateUserCreate(req); err != nil {
		api.Error(c, err)
		return
	}

	// handle
	resp, err := services.UserCreate(c, req)
	if err != nil {
		api.Error(c, err)
		return
	}
	api.Success(c, gin.H{
		"user": resp,
	})
}
