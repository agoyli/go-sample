package store

import (
	"context"
	"main/internal/models"
)

type IStore interface {
	// users
	UserFindById(ctx context.Context, id uint) (model *models.User, err error)
	UserFindBy(ctx context.Context, opts map[string]interface{}) (list *models.Users, err error)
	UserInsert(ctx context.Context, data *models.User) (model *models.User, err error)
	UserUpdate(ctx context.Context, data *models.User) (model *models.User, err error)
	UserDelete(ctx context.Context, ids []uint) (list *models.Users, err error)
}
