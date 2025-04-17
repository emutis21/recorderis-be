package adapters

import (
	"context"
	"recorderis/cmd/services/api/models"
	repository_drivers "recorderis/cmd/services/repository/ports/drivers"
)

type UserQueryerMockAdapter struct {
	ctx    context.Context
	driver repository_drivers.ForManagingUser
}

func (a *UserQueryerMockAdapter) CreateUser(ctx context.Context, user *models.CreateUserRequest) (*models.UserResponse, error) {
	panic("unimplemented")
}

func (a *UserQueryerMockAdapter) DeleteUser(ctx context.Context, id int) error {
	panic("unimplemented")
}

func (a *UserQueryerMockAdapter) GetUsers(ctx context.Context) ([]models.UserResponse, error) {
	panic("unimplemented")
}

func (a *UserQueryerMockAdapter) UpdateUser(ctx context.Context, id int, user *models.UpdateUserRequest) (*models.UserResponse, error) {
	panic("unimplemented")
}

func (a *UserQueryerMockAdapter) GetUserById(id int) (*models.UserResponse, error) {
	mockedUser := models.UserResponse{
		ID:          "1",
		DisplayName: "Esteban Mutis",
		Username:    "emutis",
		Email:       "smithmutis@gmail.com",
		AvatarURL:   "https://avatars.githubusercontent.com/u/105887678",
	}

	return &mockedUser, nil
}

func NewUserQueryerMockAdapter(ctx context.Context, driver repository_drivers.ForManagingUser) *UserQueryerMockAdapter {
	return &UserQueryerMockAdapter{
		ctx:    ctx,
		driver: driver,
	}
}
