package ports

import (
	"context"
	"recorderis/cmd/services/api/models"
)

type ForQueryingUser interface {
    GetUsers(ctx context.Context) ([]models.UserResponse, error)
    GetUserById(id int) (*models.UserResponse, error)
    CreateUser(ctx context.Context, user *models.CreateUserRequest) (*models.UserResponse, error)
    UpdateUser(ctx context.Context, id int, user *models.UpdateUserRequest) (*models.UserResponse, error)
    DeleteUser(ctx context.Context, id int) error
}
