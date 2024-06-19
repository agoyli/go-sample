package services

import (
	"context"
	"main/internal/models"
	"main/internal/store"
	"time"
)

type UserCreateDto struct {
}

type UserQueryDto struct {
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
		resp.Users = append(resp.Users, UserResponse{
			Id: v.Id,
		})
	}

	return
}

func UserCreate() {

}

func userToResponse(model models.User) (resp UserResponse) {
	return
}
