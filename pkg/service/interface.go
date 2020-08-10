package service

import (
	"context"

	"github.com/darksasori/finance/pkg/model"
)

type UserRepository interface {
	Insert(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
	FindOne(ctx context.Context, id interface{}) (*model.User, error)
}
