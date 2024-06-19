package store

import (
	"context"
	"main/internal/models"
)

type IStore interface {
	// users
	UserFindById(ctx context.Context, id uint) (model *models.User, err error)
	UserFindBy(ctx context.Context, opts map[string]interface{}) (list *models.Users, err error)
}
