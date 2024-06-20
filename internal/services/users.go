package services

import (
	"context"
	"main/internal/models"
	"main/internal/store"
	"time"
)

type UserCreateDto struct {
	FirstName  *string `json:"first_name" form:"first_name"`
	LastName   *string `json:"last_name" form:"last_name"`
	MiddleName *string `json:"middle_name" form:"middle_name"`
	Username   *string `json:"username" form:"username"`
	Password   *string `json:"password" form:"password"`
	Status     *string `json:"status" form:"status"`
	Phone      *string `json:"phone" form:"phone"`
	Email      *string `json:"email" form:"email"`
}

type UserQueryDto struct {
	Search *string `form:"search"`
	Sort   *string `form:"search"`
	Limit  int     `form:"limit"`
	Offset int     `form:"offset"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total"`
}

type UserResponse struct {
	Id              uint       `json:"id"`
	FirstName       *string    `json:"first_name"`
	LastName        *string    `json:"last_name"`
	MiddleName      *string    `json:"middle_name"`
	Username        *string    `json:"username"`
	Password        *string    `json:"password"`
	Status          *string    `json:"status"`
	Phone           *string    `json:"phone"`
	PhoneVerifiedAt *time.Time `json:"phone_verified_at"`
	Email           *string    `json:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	LastActive      *time.Time `json:"last_active"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

func UserList(ctx context.Context, query UserQueryDto) (resp *UsersResponse, err error) {
	list, err := store.Store().UserFindBy(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	resp = &UsersResponse{
		Users: []UserResponse{},
	}
	resp.Total = list.Total
	for _, v := range list.Users {
		resp.Users = append(resp.Users, userToResponse(*v))
	}

	return
}

func UserCreate(ctx context.Context, query UserCreateDto) (resp *UserResponse, err error) {

	return
}

func userToResponse(model models.User) (resp UserResponse) {
	resp = UserResponse{
		Id:              model.Id,
		FirstName:       model.FirstName,
		LastName:        model.LastName,
		MiddleName:      model.MiddleName,
		Username:        model.Username,
		Password:        model.Password,
		Status:          model.Status,
		Phone:           model.Phone,
		PhoneVerifiedAt: model.PhoneVerifiedAt,
		Email:           model.Email,
		EmailVerifiedAt: model.EmailVerifiedAt,
		LastActive:      model.LastActive,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
	}
	return resp
}
