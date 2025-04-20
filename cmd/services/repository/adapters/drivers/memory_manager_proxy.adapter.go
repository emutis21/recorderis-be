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
