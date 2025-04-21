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

/* descriptions */
func (a *MemoryAdapter) GetDescriptions(ctx context.Context, memoryID string) ([]models.DescriptionResponse, error) {
	return a.memoryRepo.GetDescriptions(ctx, memoryID)
}

func (a *MemoryAdapter) GetDescriptionByID(ctx context.Context, memoryID string, descriptionID string) (*models.DescriptionResponse, error) {
	return a.memoryRepo.GetDescriptionByID(ctx, memoryID, descriptionID)
}

func (a *MemoryAdapter) CreateDescription(ctx context.Context, memoryID string, req *models.CreateDescriptionRequest) (*models.DescriptionResponse, error) {
	return a.memoryRepo.CreateDescription(ctx, memoryID, req)
}

func (a *MemoryAdapter) UpdateDescription(ctx context.Context, memoryID string, descriptionID string, req *models.UpdateDescriptionRequest) (*models.DescriptionResponse, error) {
	return a.memoryRepo.UpdateDescription(ctx, memoryID, descriptionID, req)
}

func (a *MemoryAdapter) DeleteDescription(ctx context.Context, memoryID string, descriptionID string) error {
	return a.memoryRepo.DeleteDescription(ctx, memoryID, descriptionID)
}
