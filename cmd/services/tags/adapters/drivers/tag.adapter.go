package drivers

import (
	"context"
	"recorderis/cmd/services/tags/models"
	"recorderis/cmd/services/tags/ports/drivens"
	"recorderis/cmd/services/tags/ports/drivers"
)

var _ drivers.ForTag = (*TagAdapter)(nil)

type TagAdapter struct {
	tagRepo drivens.ForTagRepository
}

func NewTagAdapter(tagRepo drivens.ForTagRepository) *TagAdapter {
	return &TagAdapter{
		tagRepo: tagRepo,
	}
}

func (a *TagAdapter) GetTags(ctx context.Context) ([]models.TagResponse, error) {
	return a.tagRepo.GetTags(ctx)
}

func (a *TagAdapter) GetTagByID(ctx context.Context, tagID string) (*models.TagResponse, error) {
	return a.tagRepo.GetTagByID(ctx, tagID)
}

func (a *TagAdapter) CreateTag(ctx context.Context, req *models.CreateTagRequest) (*models.TagResponse, error) {
	return a.tagRepo.CreateTag(ctx, req)
}

func (a *TagAdapter) UpdateTag(ctx context.Context, tagID string, req *models.UpdateTagRequest) (*models.TagResponse, error) {
	return a.tagRepo.UpdateTag(ctx, tagID, req)
}

func (a *TagAdapter) DeleteTag(ctx context.Context, tagID string) error {
	return a.tagRepo.DeleteTag(ctx, tagID)
}

func (a *TagAdapter) AssociateMemoryWithTag(ctx context.Context, memoryID string, tagID string) error {
	return a.tagRepo.AssociateMemoryWithTag(ctx, memoryID, tagID)
}

func (a *TagAdapter) DisassociateMemoryFromTag(ctx context.Context, memoryID string, tagID string) error {
	return a.tagRepo.DisassociateMemoryFromTag(ctx, memoryID, tagID)
}

func (a *TagAdapter) GetTagsByMemoryID(ctx context.Context, memoryID string) ([]models.TagResponse, error) {
	return a.tagRepo.GetTagsByMemoryID(ctx, memoryID)
}
