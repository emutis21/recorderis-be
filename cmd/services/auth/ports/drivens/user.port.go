package drivens

import (
	"context"
	userModels "recorderis/cmd/services/api/models"
)

type ForUserRepository interface {
	CreateUser(ctx context.Context, user *userModels.CreateUserRequest) (*userModels.User, error)
	FindUserByEmail(ctx context.Context, email string) (*userModels.User, error)
	FindUserById(ctx context.Context, id string) (*userModels.User, error)
}
