package adapters

import (
	"context"
	"recorderis/cmd/services/repository"
	"recorderis/cmd/services/repository/models"
)

type UserManagerProxyAdapter struct {
	ctx        context.Context
	repository *repository.Repository
}

func (a *UserManagerProxyAdapter) GetUserById(id int) (*models.User, error) {
	return a.repository.GetUserById(a.ctx, id)
}

func (a *UserManagerProxyAdapter) FindUserByUserID(ctx context.Context, userID string) (*models.User, error) {
    return a.repository.FindUserByUserID(ctx, userID)
}

func (a *UserManagerProxyAdapter) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := a.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *UserManagerProxyAdapter) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := a.repository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *UserManagerProxyAdapter) DeleteUser(ctx context.Context, id int) error {
	return a.repository.DeleteUser(ctx, id)
}

func (a *UserManagerProxyAdapter) GetUsers(ctx context.Context) ([]models.User, error) {
	return a.repository.GetUsers(ctx)
}

func NewUserManagerProxyAdapter(ctx context.Context, repository *repository.Repository) UserManagerProxyAdapter {
	return UserManagerProxyAdapter{
		ctx:        ctx,
		repository: repository,
	}
}

func (a *UserManagerProxyAdapter) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return a.repository.FindUserByEmail(ctx, email)
}
