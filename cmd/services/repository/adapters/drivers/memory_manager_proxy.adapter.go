package adapters

import (
	"context"
	"recorderis/cmd/services/repository"
	"recorderis/cmd/services/repository/models"
)

type MemoryManagerProxyAdapter struct {
	ctx        context.Context
	repository *repository.Repository
}

func NewMemoryManagerProxyAdapter(ctx context.Context, repository *repository.Repository) *MemoryManagerProxyAdapter {
	return &MemoryManagerProxyAdapter{
		ctx:        ctx,
		repository: repository,
	}
}

func (a *MemoryManagerProxyAdapter) GetMemories(ctx context.Context, userID string) ([]models.Memory, error) {
	return a.repository.GetMemories(ctx, userID)
}

func (a *MemoryManagerProxyAdapter) GetMemoryByMemoryID(ctx context.Context, memoryID string) (*models.Memory, error) {
	return a.repository.GetMemoryByMemoryID(ctx, memoryID)
}

func (a *MemoryManagerProxyAdapter) CreateMemory(ctx context.Context, memory *models.Memory) (*models.Memory, error) {
	err := a.repository.CreateMemory(ctx, memory)
	if err != nil {
		return nil, err
	}
	return memory, nil
}

func (a *MemoryManagerProxyAdapter) UpdateMemory(ctx context.Context, memory *models.Memory) (*models.Memory, error) {
	err := a.repository.UpdateMemory(ctx, memory)
	if err != nil {
		return nil, err
	}
	return memory, nil
}

func (a *MemoryManagerProxyAdapter) DeleteMemory(ctx context.Context, memoryID string) error {
	return a.repository.DeleteMemory(ctx, memoryID)
}

/* descriptions */
func (a *MemoryManagerProxyAdapter) GetDescriptions(ctx context.Context, memoryID string) ([]models.Description, error) {
	return a.repository.GetDescriptions(ctx, memoryID)
}

func (a *MemoryManagerProxyAdapter) GetDescriptionByID(ctx context.Context, memoryID string, descriptionID string) (*models.Description, error) {
	return a.repository.GetDescriptionByID(ctx, memoryID, descriptionID)
}

func (a *MemoryManagerProxyAdapter) CreateDescription(ctx context.Context, memoryID string, description *models.Description) (*models.Description, error) {
	err := a.repository.CreateDescription(ctx, memoryID, description)
	if err != nil {
		return nil, err
	}
	return description, nil
}

func (a *MemoryManagerProxyAdapter) UpdateDescription(ctx context.Context, memoryID string, description *models.Description) (*models.Description, error) {
	err := a.repository.UpdateDescription(ctx, memoryID, description)
	if err != nil {
		return nil, err
	}
	return description, nil
}

func (a *MemoryManagerProxyAdapter) DeleteDescription(ctx context.Context, memoryID string, descriptionID string) error {
	return a.repository.DeleteDescription(ctx, memoryID, descriptionID)
}
