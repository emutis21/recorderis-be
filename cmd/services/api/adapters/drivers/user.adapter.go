package adapters

import (
	"context"
	dashboard_api "recorderis/cmd/services/api"
	api_models "recorderis/cmd/services/api/models"
)

type UserAdapter struct {
	ctx          context.Context
	dashboardApi *dashboard_api.DashboardApi
}

func (a *UserAdapter) GetUsers(ctx context.Context) ([]api_models.UserResponse, error) {
	return a.dashboardApi.GetUsers(a.ctx)
}

func (a *UserAdapter) GetUserById(id int) (*api_models.UserResponse, error) {
	return a.dashboardApi.GetUserById(a.ctx, id)
}

func (a *UserAdapter) CreateUser(ctx context.Context, user *api_models.CreateUserRequest) (*api_models.UserResponse, error) {
	return a.dashboardApi.CreateUser(a.ctx, user)
}

func (a *UserAdapter) DeleteUser(ctx context.Context, id int) error {
	return a.dashboardApi.DeleteUser(ctx, id)
}

func (a *UserAdapter) UpdateUser(ctx context.Context, id int, user *api_models.UpdateUserRequest) (*api_models.UserResponse, error) {
	return a.dashboardApi.UpdateUser(ctx, id, user)
}

func CreateUserAdapter(ctx context.Context, dashboardApi *dashboard_api.DashboardApi) *UserAdapter {
	return &UserAdapter{
		ctx:          ctx,
		dashboardApi: dashboardApi,
	}
}
