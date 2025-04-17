package ports

import (
	"context"
	"recorderis/cmd/services/repository/models"
)

type ForManagingUser interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
}
