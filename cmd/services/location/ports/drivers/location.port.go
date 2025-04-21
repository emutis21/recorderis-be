package drivers

import (
	"context"
	"recorderis/cmd/services/location/models"
)

type ForLocation interface {
	/* metodos for locations */
	GetLocations(ctx context.Context) ([]models.LocationResponse, error)
	GetLocationByID(ctx context.Context, locationID string) (*models.LocationResponse, error)
	CreateLocation(ctx context.Context, req *models.CreateLocationRequest) (*models.LocationResponse, error)
	UpdateLocation(ctx context.Context, locationID string, req *models.UpdateLocationRequest) (*models.LocationResponse, error)
	DeleteLocation(ctx context.Context, locationID string) error

	/* relationships with memories */
	AssociateMemoryWithLocation(ctx context.Context, memoryID string, locationID string) error
	DisassociateMemoryFromLocation(ctx context.Context, memoryID string, locationID string) error
	GetLocationsByMemoryID(ctx context.Context, memoryID string) ([]models.LocationResponse, error)
}
