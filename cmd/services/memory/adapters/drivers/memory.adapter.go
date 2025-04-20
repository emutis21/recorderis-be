package drivers

import (
	"context"
	"recorderis/cmd/services/memory/models"
	"recorderis/cmd/services/memory/ports/drivens"
	"recorderis/cmd/services/memory/ports/drivers"
)

type MemoryAdapter struct {
	memoryRepo drivens.ForMemoryRepository
}

func NewMemoryAdapter(memoryRepo drivens.ForMemoryRepository) drivers.ForMemory {
	return &MemoryAdapter{
		memoryRepo: memoryRepo,
	}
}

func (a *MemoryAdapter) GetMemories(ctx context.Context, userID string) ([]models.MemoryResponse, error) {
	return a.memoryRepo.GetMemories(ctx, userID)
}

func (a *MemoryAdapter) GetMemoryByID(ctx context.Context, memoryID string) (*models.MemoryResponse, error) {
	return a.memoryRepo.GetMemoryByID(ctx, memoryID)
}

type contextKey string

func (a *MemoryAdapter) CreateMemory(ctx context.Context, req *models.CreateMemoryRequest) (*models.MemoryResponse, error) {
	const userIDKey contextKey = "userID"

	ctx = context.WithValue(ctx, userIDKey, ctx.Value(userIDKey))
	return a.memoryRepo.CreateMemory(ctx, req)
}

func (a *MemoryAdapter) UpdateMemory(ctx context.Context, memoryID string, req *models.UpdateMemoryRequest) (*models.MemoryResponse, error) {
	return a.memoryRepo.UpdateMemory(ctx, memoryID, req)
}

func (a *MemoryAdapter) DeleteMemory(ctx context.Context, memoryID string) error {
	return a.memoryRepo.DeleteMemory(ctx, memoryID)
}
