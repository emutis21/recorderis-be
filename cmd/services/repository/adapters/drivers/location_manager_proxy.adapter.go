package adapters

import (
	"context"
	"recorderis/cmd/services/repository"
	"recorderis/cmd/services/repository/models"
)

type LocationManagerProxyAdapter struct {
	ctx        context.Context
	repository *repository.Repository
}

func NewLocationManagerProxyAdapter(ctx context.Context, repository *repository.Repository) *LocationManagerProxyAdapter {
	return &LocationManagerProxyAdapter{
		ctx:        ctx,
		repository: repository,
	}
}

/* crud basic operations */
func (a *LocationManagerProxyAdapter) GetLocations(ctx context.Context) ([]models.Location, error) {
	return a.repository.GetLocations(ctx)
}

func (a *LocationManagerProxyAdapter) GetLocationByID(ctx context.Context, locationID string) (*models.Location, error) {
	return a.repository.GetLocationByID(ctx, locationID)
}

func (a *LocationManagerProxyAdapter) CreateLocation(ctx context.Context, location *models.Location) (*models.Location, error) {
	err := a.repository.CreateLocation(ctx, location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (a *LocationManagerProxyAdapter) UpdateLocation(ctx context.Context, location *models.Location) (*models.Location, error) {
	err := a.repository.UpdateLocation(ctx, location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (a *LocationManagerProxyAdapter) DeleteLocation(ctx context.Context, locationID string) error {
	return a.repository.DeleteLocation(ctx, locationID)
}

/* relationships with memories */
func (a *LocationManagerProxyAdapter) AssociateMemoryWithLocation(ctx context.Context, memoryID string, locationID string) error {
	return a.repository.AssociateMemoryWithLocation(ctx, memoryID, locationID)
}

func (a *LocationManagerProxyAdapter) DisassociateMemoryFromLocation(ctx context.Context, memoryID string, locationID string) error {
	return a.repository.DisassociateMemoryFromLocation(ctx, memoryID, locationID)
}

func (a *LocationManagerProxyAdapter) GetLocationsByMemoryID(ctx context.Context, memoryID string) ([]models.Location, error) {
	return a.repository.GetLocationsByMemoryID(ctx, memoryID)
}

func (a *LocationManagerProxyAdapter) GetMemoriesByLocationID(ctx context.Context, locationID string) ([]string, error) {
	return a.repository.GetMemoriesByLocationID(ctx, locationID)
}
