package models

import "time"

type User struct {
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

type Users struct {
	Users []*User
	Total int
}
