package adapters

import (
	"context"
	api_models "recorderis/cmd/services/api/models"
	repository_models "recorderis/cmd/services/repository/models"
	repository_drivers "recorderis/cmd/services/repository/ports/drivers"
	"strconv"
)

type (
	GetUserById func(id int) (string, error)
)

type UserQueryerAdapter struct {
	ctx    context.Context
	driver repository_drivers.ForManagingUser
}

func (a *UserQueryerAdapter) GetUsers(ctx context.Context) ([]api_models.UserResponse, error) {
	users, err := a.driver.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var userResponses []api_models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, api_models.UserResponse{
			ID:          user.UserID,
			Username:    user.Username,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			AvatarURL:   user.AvatarURL,
		})
	}

	return userResponses, nil
}

func (a *UserQueryerAdapter) GetUserById(id int) (*api_models.UserResponse, error) {
	user, err := a.driver.GetUserById(id)

	if err != nil {
		return nil, err
	}

	mappedUser := api_models.UserResponse{
		ID:          strconv.Itoa(user.ID),
		UserID:      user.UserID,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		AvatarURL:   user.AvatarURL,
	}

	return &mappedUser, nil
}

func (a *UserQueryerAdapter) CreateUser(ctx context.Context, user *api_models.CreateUserRequest) (*api_models.UserResponse, error) {
	createdUser, err := a.driver.CreateUser(ctx, &repository_models.User{
		Username:     user.Username,
		DisplayName:  user.DisplayName,
		Email:        user.Email,
		PasswordHash: user.Password,
	})

	if err != nil {
		return nil, err
	}

	return &api_models.UserResponse{
		ID:          strconv.Itoa(createdUser.ID),
		UserID:      createdUser.UserID,
		Username:    createdUser.Username,
		DisplayName: createdUser.DisplayName,
		Email:       createdUser.Email,
		AvatarURL:   createdUser.AvatarURL,
	}, nil
}

func (a *UserQueryerAdapter) DeleteUser(ctx context.Context, id int) error {
	return a.driver.DeleteUser(ctx, id)
}

func (a *UserQueryerAdapter) UpdateUser(ctx context.Context, id int, user *api_models.UpdateUserRequest) (*api_models.UserResponse, error) {

	existingUser, err := a.driver.GetUserById(id)
	if err != nil {
		return nil, err
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.DisplayName != "" {
		existingUser.DisplayName = user.DisplayName
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}

	updatedUser, err := a.driver.UpdateUser(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	return &api_models.UserResponse{
		ID:          strconv.Itoa(updatedUser.ID),
		UserID:      updatedUser.UserID,
		Username:    updatedUser.Username,
		DisplayName: updatedUser.DisplayName,
		Email:       updatedUser.Email,
		AvatarURL:   updatedUser.AvatarURL,
	}, nil
}

func NewUserQueryerAdapter(ctx context.Context, driver repository_drivers.ForManagingUser) *UserQueryerAdapter {
	return &UserQueryerAdapter{
		ctx:    ctx,
		driver: driver,
	}
}
