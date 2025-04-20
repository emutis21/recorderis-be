package drivens

import (
	"context"
	"recorderis/cmd/services/memory/models"
)

type ForMemoryRepository interface {
	GetMemories(ctx context.Context, userID string) ([]models.MemoryResponse, error)
	GetMemoryByID(ctx context.Context, memoryID string) (*models.MemoryResponse, error)
	CreateMemory(ctx context.Context, req *models.CreateMemoryRequest) (*models.MemoryResponse, error)
	UpdateMemory(ctx context.Context, memoryID string, req *models.UpdateMemoryRequest) (*models.MemoryResponse, error)
	DeleteMemory(ctx context.Context, memoryID string) error
}
