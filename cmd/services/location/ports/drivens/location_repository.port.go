package drivens

import (
	"context"
	"recorderis/cmd/services/location/models"
)

type ForLocationRepository interface {
	GetLocations(ctx context.Context) ([]models.LocationResponse, error)
	GetLocationByID(ctx context.Context, locationID string) (*models.LocationResponse, error)
	CreateLocation(ctx context.Context, req *models.CreateLocationRequest) (*models.LocationResponse, error)
	UpdateLocation(ctx context.Context, locationID string, req *models.UpdateLocationRequest) (*models.LocationResponse, error)
	DeleteLocation(ctx context.Context, locationID string) error

	AssociateMemoryWithLocation(ctx context.Context, memoryID string, locationID string) error
	DisassociateMemoryFromLocation(ctx context.Context, memoryID string, locationID string) error
	GetLocationsByMemoryID(ctx context.Context, memoryID string) ([]models.LocationResponse, error)
	GetMemoriesByLocationID(ctx context.Context, locationID string) ([]string, error)
}
