package models

import "time"

type User struct {
	Id              uint
	FirstName       *string
	LastName        *string
	MiddleName      *string
	Username        *string
	Password        *string
	Status          *string
	Phone           *string
	PhoneVerifiedAt *time.Time
	Email           *string
	EmailVerifiedAt *time.Time
	LastActive      *time.Time
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

type Users struct {
	Users []*User
	Total int
}
