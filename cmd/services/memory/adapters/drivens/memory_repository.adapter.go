package drivens

import (
	"context"
	"errors"
	"recorderis/cmd/services/memory/models"
	"recorderis/cmd/services/memory/ports/drivens"
	repo_adapters "recorderis/cmd/services/repository/adapters/drivers"
	repo_models "recorderis/cmd/services/repository/models"
	"recorderis/internals/constants"
	"recorderis/internals/utils"
	"strconv"
)

type MemoryRepositoryAdapter struct {
	memoryRepo *repo_adapters.MemoryManagerProxyAdapter
}

func NewMemoryRepositoryAdapter(memoryRepo *repo_adapters.MemoryManagerProxyAdapter) drivens.ForMemoryRepository {
	return &MemoryRepositoryAdapter{
		memoryRepo: memoryRepo,
	}
}

func (a *MemoryRepositoryAdapter) GetMemories(ctx context.Context, userID string) ([]models.MemoryResponse, error) {
	memories, err := a.memoryRepo.GetMemories(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []models.MemoryResponse
	for _, memory := range memories {
		responses = append(responses, mapMemoryToResponse(&memory))
	}

	return responses, nil
}

func (a *MemoryRepositoryAdapter) GetMemoryByID(ctx context.Context, memoryID string) (*models.MemoryResponse, error) {
	memory, err := a.memoryRepo.GetMemoryByMemoryID(ctx, memoryID)
	if err != nil {
		return nil, err
	}

	response := mapMemoryToResponse(memory)
	return &response, nil
}

func (a *MemoryRepositoryAdapter) CreateMemory(ctx context.Context, req *models.CreateMemoryRequest) (*models.MemoryResponse, error) {
	userIDValue := ctx.Value(constants.UserIDKey)
	if userIDValue == nil {
		return nil, errors.New("user not authenticated")
	}

	userID, ok := userIDValue.(string)
	if !ok {
		return nil, errors.New("invalid user ID format")
	}

	memory := &repo_models.Memory{
		UserID:   userID,
		Title:    req.Title,
		Date:     req.Date.Time(),
		IsPublic: req.IsPublic,
	}

	if len(req.Descriptions) > 0 {
		for _, desc := range req.Descriptions {
			memory.Descriptions = append(memory.Descriptions, &repo_models.Description{
				Text:  desc.Text,
				Index: desc.Index,
			})
		}
	}

	createdMemory, err := a.memoryRepo.CreateMemory(ctx, memory)
	if err != nil {
		return nil, err
	}

	response := mapMemoryToResponse(createdMemory)
	return &response, nil
}

func (a *MemoryRepositoryAdapter) UpdateMemory(ctx context.Context, memoryID string, req *models.UpdateMemoryRequest) (*models.MemoryResponse, error) {
	existingMemory, err := a.memoryRepo.GetMemoryByMemoryID(ctx, memoryID)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		existingMemory.Title = req.Title
	}

	if req.Date != (utils.JSONTime{}) {
		existingMemory.Date = req.Date.Time()
	}

	existingMemory.IsPublicPtr = req.IsPublic

	if req.Index != nil {
		existingMemory.IndexPtr = req.Index
	}

	updatedMemory, err := a.memoryRepo.UpdateMemory(ctx, existingMemory)
	if err != nil {
		return nil, err
	}

	response := mapMemoryToResponse(updatedMemory)

	return &response, nil
}

func (a *MemoryRepositoryAdapter) DeleteMemory(ctx context.Context, memoryID string) error {
	return a.memoryRepo.DeleteMemory(ctx, memoryID)
}

func mapMemoryToResponse(memory *repo_models.Memory) models.MemoryResponse {
	response := models.MemoryResponse{
		ID:        strconv.Itoa(memory.ID),
		MemoryID:  memory.MemoryID,
		Title:     memory.Title,
		Date:      memory.Date,
		IsPublic:  memory.IsPublic,
		CreatedAt: memory.CreatedAt,
		UpdatedAt: memory.UpdatedAt,
	}

	if memory.Descriptions != nil {
		for _, desc := range memory.Descriptions {
			response.Descriptions = append(response.Descriptions, mapDescriptionToResponse(desc))
		}
	}

	if memory.Photos != nil {
		for _, photo := range memory.Photos {
			response.Photos = append(response.Photos, models.PhotoResponse{
				ID:       strconv.Itoa(photo.ID),
				PhotoID:  photo.PhotoID,
				FileName: photo.FileName,
				URL:      photo.URL,
				Index:    photo.Index,
			})
		}
	}

	return response
}

func (a *MemoryRepositoryAdapter) GetDescriptions(ctx context.Context, memoryID string) ([]models.DescriptionResponse, error) {
	descriptions, err := a.memoryRepo.GetDescriptions(ctx, memoryID)
	if err != nil {
		return nil, err
	}

	responses := make([]models.DescriptionResponse, 0, len(descriptions))
	for _, desc := range descriptions {
		responses = append(responses, mapDescriptionToResponse(&desc))
	}

	return responses, nil
}

func (a *MemoryRepositoryAdapter) GetDescriptionByID(ctx context.Context, memoryID string, descriptionID string) (*models.DescriptionResponse, error) {
	description, err := a.memoryRepo.GetDescriptionByID(ctx, memoryID, descriptionID)
	if err != nil {
		return nil, err
	}

	response := mapDescriptionToResponse(description)
	return &response, nil
}

func (a *MemoryRepositoryAdapter) CreateDescription(ctx context.Context, memoryID string, req *models.CreateDescriptionRequest) (*models.DescriptionResponse, error) {
	_, err := a.memoryRepo.GetMemoryByMemoryID(ctx, memoryID)
	if err != nil {
		return nil, err
	}

	description := &repo_models.Description{
		MemoryID: memoryID,
		Text:     req.Text,
		Index:    req.Index,
	}

	createdDescription, err := a.memoryRepo.CreateDescription(ctx, memoryID, description)
	if err != nil {
		return nil, err
	}

	response := mapDescriptionToResponse(createdDescription)
	return &response, nil
}

func (a *MemoryRepositoryAdapter) UpdateDescription(ctx context.Context, memoryID string, descriptionID string, req *models.UpdateDescriptionRequest) (*models.DescriptionResponse, error) {
	description, err := a.memoryRepo.GetDescriptionByID(ctx, memoryID, descriptionID)
	if err != nil {
		return nil, err
	}

	if req.Text != "" {
		description.Text = req.Text
	}

	if req.Index != nil {
		description.Index = *req.Index
	}

	description.DescriptionID = descriptionID

	updatedDescription, err := a.memoryRepo.UpdateDescription(ctx, memoryID, description)
	if err != nil {
		return nil, err
	}

	response := mapDescriptionToResponse(updatedDescription)
	return &response, nil
}

func (a *MemoryRepositoryAdapter) DeleteDescription(ctx context.Context, memoryID string, descriptionID string) error {
	return a.memoryRepo.DeleteDescription(ctx, memoryID, descriptionID)
}

func mapDescriptionToResponse(description *repo_models.Description) models.DescriptionResponse {
	return models.DescriptionResponse{
		ID:            strconv.Itoa(description.ID),
		DescriptionID: description.DescriptionID,
		Text:          description.Text,
		Index:         description.Index,
		Version:       description.Version,
	}
}
