package ports

import (
	"context"
	"recorderis/cmd/services/api/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}