package drivens

import (
	"context"
	"recorderis/cmd/services/location/models"
	"recorderis/cmd/services/location/ports/drivens"
	repo_adapters "recorderis/cmd/services/repository/adapters/drivers"
	repo_models "recorderis/cmd/services/repository/models"
	"strconv"
)

var _ drivens.ForLocationRepository = (*LocationRepositoryAdapter)(nil)

type LocationRepositoryAdapter struct {
	locationRepo *repo_adapters.LocationManagerProxyAdapter
}

func NewLocationRepositoryAdapter(locationRepo *repo_adapters.LocationManagerProxyAdapter) *LocationRepositoryAdapter {
	return &LocationRepositoryAdapter{
		locationRepo: locationRepo,
	}
}

func (a *LocationRepositoryAdapter) GetLocations(ctx context.Context) ([]models.LocationResponse, error) {
	locations, err := a.locationRepo.GetLocations(ctx)
	if err != nil {
		return nil, err
	}

	var response []models.LocationResponse
	for _, loc := range locations {
		response = append(response, mapLocationToResponse(&loc))
	}

	return response, nil
}

func (a *LocationRepositoryAdapter) GetLocationByID(ctx context.Context, locationID string) (*models.LocationResponse, error) {
	location, err := a.locationRepo.GetLocationByID(ctx, locationID)
	if err != nil {
		return nil, err
	}

	response := mapLocationToResponse(location)
	return &response, nil
}

func (a *LocationRepositoryAdapter) CreateLocation(ctx context.Context, req *models.CreateLocationRequest) (*models.LocationResponse, error) {
	location := &repo_models.Location{
		Location:  req.Location,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
		City:      req.City,
		Country:   req.Country,
	}

	createdLocation, err := a.locationRepo.CreateLocation(ctx, location)
	if err != nil {
		return nil, err
	}

	response := mapLocationToResponse(createdLocation)
	return &response, nil
}

func (a *LocationRepositoryAdapter) UpdateLocation(ctx context.Context, locationID string, req *models.UpdateLocationRequest) (*models.LocationResponse, error) {
	location, err := a.locationRepo.GetLocationByID(ctx, locationID)
	if err != nil {
		return nil, err
	}

	if req.Location != "" {
		location.Location = req.Location
	}
	/* TODO: improve this, we should not update longitude and latitude if they are not provided */
	if req.Longitude != nil {
		location.Longitude = *req.Longitude
	}
	if req.Latitude != nil {
		location.Latitude = *req.Latitude
	}
	if req.City != "" {
		location.City = req.City
	}
	if req.Country != "" {
		location.Country = req.Country
	}

	updatedLocation, err := a.locationRepo.UpdateLocation(ctx, location)
	if err != nil {
		return nil, err
	}

	response := mapLocationToResponse(updatedLocation)
	return &response, nil
}

func (a *LocationRepositoryAdapter) DeleteLocation(ctx context.Context, locationID string) error {
	return a.locationRepo.DeleteLocation(ctx, locationID)
}

func (a *LocationRepositoryAdapter) AssociateMemoryWithLocation(ctx context.Context, memoryID string, locationID string) error {
	return a.locationRepo.AssociateMemoryWithLocation(ctx, memoryID, locationID)
}

func (a *LocationRepositoryAdapter) DisassociateMemoryFromLocation(ctx context.Context, memoryID string, locationID string) error {
	return a.locationRepo.DisassociateMemoryFromLocation(ctx, memoryID, locationID)
}

func (a *LocationRepositoryAdapter) GetLocationsByMemoryID(ctx context.Context, memoryID string) ([]models.LocationResponse, error) {
	locations, err := a.locationRepo.GetLocationsByMemoryID(ctx, memoryID)
	if err != nil {
		return nil, err
	}

	var response []models.LocationResponse
	for _, loc := range locations {
		response = append(response, mapLocationToResponse(&loc))
	}

	return response, nil
}

func (a *LocationRepositoryAdapter) GetMemoriesByLocationID(ctx context.Context, locationID string) ([]string, error) {
	return a.locationRepo.GetMemoriesByLocationID(ctx, locationID)
}

func mapLocationToResponse(location *repo_models.Location) models.LocationResponse {
	return models.LocationResponse{
		ID:         strconv.Itoa(location.ID),
		LocationID: location.LocationID,
		Location:   location.Location,
		Longitude:  location.Longitude,
		Latitude:   location.Latitude,
		City:       location.City,
		Country:    location.Country,
	}
}
