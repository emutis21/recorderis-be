package main

import (
	"context"
	"log/slog"
	"recorderis/internals/config"

	api_service "recorderis/cmd/services/api"
	api_driven_adapters "recorderis/cmd/services/api/adapters/drivens"
	api_driver_adapters "recorderis/cmd/services/api/adapters/drivers"
	api_ports "recorderis/cmd/services/api/ports/drivers"

	auth_driven_adapters "recorderis/cmd/services/auth/adapters/drivens"
	auth_driver_adapters "recorderis/cmd/services/auth/adapters/drivers"
	auth_models "recorderis/cmd/services/auth/models"
	auth_driven_ports "recorderis/cmd/services/auth/ports/drivens"
	auth_ports "recorderis/cmd/services/auth/ports/drivers"

	memory_driven_adapters "recorderis/cmd/services/memory/adapters/drivens"
	memory_driver_adapters "recorderis/cmd/services/memory/adapters/drivers"
	memory_ports "recorderis/cmd/services/memory/ports/drivers"

	"recorderis/cmd/services/repository"
	repository_adapters "recorderis/cmd/services/repository/adapters/drivers"

	location_driven_adapters "recorderis/cmd/services/location/adapters/drivens"
	location_driver_adapters "recorderis/cmd/services/location/adapters/drivers"
	location_ports "recorderis/cmd/services/location/ports/drivers"
)

func Compose() (api_ports.ForUser, auth_ports.ForAuth, memory_ports.ForMemory, auth_driven_ports.ForTokenManager, location_ports.ForLocation, error) {
	ctx := context.Background()
	cfg := config.LoadConfig()

	// Create (persistence) repository with error handling
	repo, err := repository.NewRepository()
	if err != nil {
		slog.Error("Failed to create repository", "error", err)

		return nil, nil, nil, nil, nil, err
	}

	// Create repository drivers
	userManagerProxyAdapter := repository_adapters.NewUserManagerProxyAdapter(ctx, repo)

	// Create dashboard api drivens
	userQueryerAdapter := api_driven_adapters.NewUserQueryerAdapter(ctx, &userManagerProxyAdapter)

	// Create dashboard api
	dashboardApi := api_service.NewDashboardApi(userQueryerAdapter)

	// Create dashboard api drivers
	userAdapter := api_driver_adapters.CreateUserAdapter(ctx, dashboardApi)

	// Create auth api
	passwordMgr := auth_driven_adapters.NewBcryptAdapter()

	jwtConfig := auth_models.TokenConfig{
		AccessTokenDuration:  auth_models.AccessTokenDuration,
		RefreshTokenDuration: auth_models.DefaultRefreshTokenDuration,
		SigningKey:           []byte(cfg.JWTSecret),
		Issuer:               "recorderis-api",
	}

	tokenMgr := auth_driven_adapters.NewJWTAdapter(jwtConfig)

	tokenRepo := auth_driven_adapters.NewGormTokenRepository(repo.GetDB())

	userRepoAdapter := auth_driven_adapters.NewUserRepositoryAdapter(&userManagerProxyAdapter)

	authAdapter := auth_driver_adapters.NewAuthAdapter(
		userRepoAdapter,
		tokenMgr,
		tokenRepo,
		passwordMgr,
	)

	memoryManagerProxyAdapter := repository_adapters.NewMemoryManagerProxyAdapter(ctx, repo)
	memoryRepoAdapter := memory_driven_adapters.NewMemoryRepositoryAdapter(memoryManagerProxyAdapter)
	memoryAdapter := memory_driver_adapters.NewMemoryAdapter(memoryRepoAdapter)

	locationManagerProxyAdapter := repository_adapters.NewLocationManagerProxyAdapter(ctx, repo)
	locationRepoAdapter := location_driven_adapters.NewLocationRepositoryAdapter(locationManagerProxyAdapter)
	locationAdapter := location_driver_adapters.NewLocationAdapter(locationRepoAdapter)

	return userAdapter, authAdapter, memoryAdapter, tokenMgr, locationAdapter, nil
}

func ComposeMock() (api_ports.ForUser, error) {
	ctx := context.Background()

	repo, err := repository.NewRepository()
	if err != nil {
		slog.Error("Failed to create mock repository", "error", err)

		return nil, err
	}

	userManagerProxyAdapter := repository_adapters.NewUserManagerProxyAdapter(ctx, repo)

	userQueryerAdapter := api_driven_adapters.NewUserQueryerMockAdapter(ctx, &userManagerProxyAdapter)

	dashboardApi := api_service.NewDashboardApi(userQueryerAdapter)

	userAdapter := api_driver_adapters.CreateUserAdapter(ctx, dashboardApi)

	return userAdapter, nil
}
