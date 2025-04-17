package dashboard_api

import (
	"context"
	api_models "recorderis/cmd/services/api/models"
	ports "recorderis/cmd/services/api/ports/drivens"
)

type DashboardApi struct {
	// drivens
	userQueryer ports.ForQueryingUser
}

func (d *DashboardApi) GetUsers(ctx context.Context) ([]api_models.UserResponse, error) {
	return d.userQueryer.GetUsers(ctx)
}

func (d *DashboardApi) GetUserById(ctx context.Context, id int) (*api_models.UserResponse, error) {
	return d.userQueryer.GetUserById(id)
}

func (d *DashboardApi) CreateUser(ctx context.Context, user *api_models.CreateUserRequest) (*api_models.UserResponse, error) {
	return d.userQueryer.CreateUser(ctx, user)
}

func (d *DashboardApi) DeleteUser(ctx context.Context, id int) error {
	return d.userQueryer.DeleteUser(ctx, id)
}

func (d *DashboardApi) UpdateUser(ctx context.Context, id int, user *api_models.UpdateUserRequest) (*api_models.UserResponse, error) {
	return d.userQueryer.UpdateUser(ctx, id, user)
}

func NewDashboardApi(userQueryer ports.ForQueryingUser) *DashboardApi {
	return &DashboardApi{userQueryer: userQueryer}
}
