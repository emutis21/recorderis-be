package drivens

import (
	"context"
	api_models "recorderis/cmd/services/api/models"
	"recorderis/cmd/services/auth/ports/drivens"
	repo_adapters "recorderis/cmd/services/repository/adapters/drivers"
	repo_models "recorderis/cmd/services/repository/models"
)

var _ drivens.ForUserRepository = (*UserRepositoryAdapter)(nil)

type UserRepositoryAdapter struct {
	userRepo *repo_adapters.UserManagerProxyAdapter
}

func NewUserRepositoryAdapter(userRepo *repo_adapters.UserManagerProxyAdapter) *UserRepositoryAdapter {
	return &UserRepositoryAdapter{
		userRepo: userRepo,
	}
}

func (a *UserRepositoryAdapter) CreateUser(ctx context.Context, req *api_models.CreateUserRequest) (*api_models.User, error) {
	repoUser := &repo_models.User{
		Username:     req.Username,
		DisplayName:  req.DisplayName,
		Email:        req.Email,
		PasswordHash: req.Password,
	}

	createdUser, err := a.userRepo.CreateUser(ctx, repoUser)
	if err != nil {
		return nil, err
	}

	return &api_models.User{
		ID:           createdUser.ID,
		UserID:       createdUser.UserID,
		Username:     createdUser.Username,
		DisplayName:  createdUser.DisplayName,
		AvatarURL:    createdUser.AvatarURL,
		Email:        createdUser.Email,
		PasswordHash: createdUser.PasswordHash,
		Role:         string(createdUser.Role),
	}, nil
}

func (a *UserRepositoryAdapter) FindUserByEmail(ctx context.Context, email string) (*api_models.User, error) {
	repoUser, err := a.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &api_models.User{
		ID:           repoUser.ID,
		UserID:       repoUser.UserID,
		Username:     repoUser.Username,
		DisplayName:  repoUser.DisplayName,
		AvatarURL:    repoUser.AvatarURL,
		Email:        repoUser.Email,
		PasswordHash: repoUser.PasswordHash,
		Role:         string(repoUser.Role),
	}, nil
}

func (a *UserRepositoryAdapter) FindUserById(ctx context.Context, id string) (*api_models.User, error) {
	repoUser, err := a.userRepo.FindUserByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &api_models.User{
		ID:           repoUser.ID,
		UserID:       repoUser.UserID,
		Username:     repoUser.Username,
		DisplayName:  repoUser.DisplayName,
		AvatarURL:    repoUser.AvatarURL,
		Email:        repoUser.Email,
		PasswordHash: repoUser.PasswordHash,
		Role:         string(repoUser.Role),
	}, nil
}
