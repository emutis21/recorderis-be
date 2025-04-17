// cmd/services/auth/adapters/drivens/user_repository.adapter.go
package drivens

import (
	"context"
	"fmt"
	api_models "recorderis/cmd/services/api/models"
	"recorderis/cmd/services/auth/ports/drivens"
	repo_adapters "recorderis/cmd/services/repository/adapters/drivers"
	repo_models "recorderis/cmd/services/repository/models"
	"strconv"
)

// Verificar que implementa la interfaz
var _ drivens.ForUserRepository = (*UserRepositoryAdapter)(nil)

// UserRepositoryAdapter adapta el repositorio a la interfaz que espera auth
type UserRepositoryAdapter struct {
	userRepo *repo_adapters.UserManagerProxyAdapter
}

func NewUserRepositoryAdapter(userRepo *repo_adapters.UserManagerProxyAdapter) *UserRepositoryAdapter {
	return &UserRepositoryAdapter{
		userRepo: userRepo,
	}
}

// CreateUser convierte el request de API y llama al repositorio subyacente
func (a *UserRepositoryAdapter) CreateUser(ctx context.Context, req *api_models.CreateUserRequest) (*api_models.User, error) {
	repoUser := &repo_models.User{
		Username:     req.Username,
		DisplayName:  req.DisplayName,
		Email:        req.Email,
		PasswordHash: req.Password, // Será hasheado por el servicio antes de llegar aquí
	}

	// Llamar al método del repositorio
	createdUser, err := a.userRepo.CreateUser(ctx, repoUser)
	if err != nil {
		return nil, err
	}

	// Convertir de repo_models.User a api_models.User
	return &api_models.User{
		ID:           createdUser.ID,
		UserID:       createdUser.UserID,
		Username:     createdUser.Username,
		DisplayName:  createdUser.DisplayName,
		AvatarURL:    createdUser.AvatarURL,
		Email:        createdUser.Email,
		PasswordHash: createdUser.PasswordHash, // Se incluye para verificaciones internas
		Role:         string(createdUser.Role),
	}, nil
}

func (a *UserRepositoryAdapter) FindUserByEmail(ctx context.Context, email string) (*api_models.User, error) {
	repoUser, err := a.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Convertir el resultado
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

// FindUserById busca un usuario por ID y convierte el resultado
func (a *UserRepositoryAdapter) FindUserById(ctx context.Context, id string) (*api_models.User, error) {
	// Convertir string a int para llamar al repositorio
	numericId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	// Llamar al repositorio
	repoUser, err := a.userRepo.GetUserById(numericId)
	if err != nil {
		return nil, err
	}

	// Convertir el resultado
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
