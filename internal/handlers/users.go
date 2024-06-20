package handlers

import (
	"main/internal/handlers/api"
	"main/internal/services"
	"main/internal/services/validator"
	"main/internal/utils"
	"strconv"

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

func UserUpdateHandler(c *gin.Context) {
	// bind request
	req := services.UserCreateDto{}
	if err := BindAny(c, &req); err != nil {
		api.Error(c, err)
		return
	}
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	// validate
	if err := validator.ValidateUserCreate(req); err != nil {
		api.Error(c, err)
		return
	}
	if id == 0 {
		api.Error(c, utils.ErrRequired.SetKey("id"))
		return
	}

	// handle
	resp, err := services.UserUpdate(c, uint(id), req)
	if err != nil {
		api.Error(c, err)
		return
	}
	api.Success(c, gin.H{
		"user": resp,
	})
}

func UserDeleteHandler(c *gin.Context) {
	// bind request
	idsStr := c.QueryArray("ids")
	ids := []uint{}
	for _, v := range idsStr {
		id, _ := strconv.Atoi(v)
		ids = append(ids, uint(id))
	}
	// validate
	if len(ids) == 0 {
		api.Error(c, utils.ErrRequired.SetKey("id"))
		return
	}
	// handle
	resp, err := services.UserDelete(c, ids)
	if err != nil {
		api.Error(c, err)
		return
	}
	api.Success(c, gin.H{
		"users": resp.Users,
	})
}
