package drivers

import (
	"context"
	"recorderis/cmd/services/location/models"
	"recorderis/cmd/services/location/ports/drivens"
	"recorderis/cmd/services/location/ports/drivers"
)

var _ drivers.ForLocation = (*LocationAdapter)(nil)

type LocationAdapter struct {
	locationRepo drivens.ForLocationRepository
}

func NewLocationAdapter(locationRepo drivens.ForLocationRepository) *LocationAdapter {
	return &LocationAdapter{
		locationRepo: locationRepo,
	}
}

func (a *LocationAdapter) GetLocations(ctx context.Context) ([]models.LocationResponse, error) {
	return a.locationRepo.GetLocations(ctx)
}

func (a *LocationAdapter) GetLocationByID(ctx context.Context, locationID string) (*models.LocationResponse, error) {
	return a.locationRepo.GetLocationByID(ctx, locationID)
}

func (a *LocationAdapter) CreateLocation(ctx context.Context, req *models.CreateLocationRequest) (*models.LocationResponse, error) {
	return a.locationRepo.CreateLocation(ctx, req)
}

func (a *LocationAdapter) UpdateLocation(ctx context.Context, locationID string, req *models.UpdateLocationRequest) (*models.LocationResponse, error) {
	return a.locationRepo.UpdateLocation(ctx, locationID, req)
}

func (a *LocationAdapter) DeleteLocation(ctx context.Context, locationID string) error {
	return a.locationRepo.DeleteLocation(ctx, locationID)
}

func (a *LocationAdapter) AssociateMemoryWithLocation(ctx context.Context, memoryID string, locationID string) error {
	return a.locationRepo.AssociateMemoryWithLocation(ctx, memoryID, locationID)
}

func (a *LocationAdapter) DisassociateMemoryFromLocation(ctx context.Context, memoryID string, locationID string) error {
	return a.locationRepo.DisassociateMemoryFromLocation(ctx, memoryID, locationID)
}

func (a *LocationAdapter) GetLocationsByMemoryID(ctx context.Context, memoryID string) ([]models.LocationResponse, error) {
	return a.locationRepo.GetLocationsByMemoryID(ctx, memoryID)
}
