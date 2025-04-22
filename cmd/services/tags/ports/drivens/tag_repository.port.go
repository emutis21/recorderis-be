package drivens

import (
	"context"
	"recorderis/cmd/services/tags/models"
)

type ForTagRepository interface {
	/* basic methods for Tags */
	GetTags(ctx context.Context) ([]models.TagResponse, error)
	GetTagByID(ctx context.Context, tagID string) (*models.TagResponse, error)
	CreateTag(ctx context.Context, req *models.CreateTagRequest) (*models.TagResponse, error)
	UpdateTag(ctx context.Context, tagID string, req *models.UpdateTagRequest) (*models.TagResponse, error)
	DeleteTag(ctx context.Context, tagID string) error

	/* relationships with memories */
	AssociateMemoryWithTag(ctx context.Context, memoryID string, tagID string) error
	DisassociateMemoryFromTag(ctx context.Context, memoryID string, tagID string) error
	GetTagsByMemoryID(ctx context.Context, memoryID string) ([]models.TagResponse, error)
	GetMemoriesByTagID(ctx context.Context, tagID string) ([]string, error)
}