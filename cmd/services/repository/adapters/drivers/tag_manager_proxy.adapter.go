package adapters

import (
	"context"
	"recorderis/cmd/services/repository"
	"recorderis/cmd/services/repository/models"
)

type TagManagerProxyAdapter struct {
	ctx        context.Context
	repository *repository.Repository
}

func NewTagManagerProxyAdapter(ctx context.Context, repository *repository.Repository) *TagManagerProxyAdapter {
	return &TagManagerProxyAdapter{
		ctx:        ctx,
		repository: repository,
	}
}

func (a *TagManagerProxyAdapter) GetTags(ctx context.Context) ([]models.Tag, error) {
	return a.repository.GetTags(ctx)
}

func (a *TagManagerProxyAdapter) GetTagByID(ctx context.Context, tagID string) (*models.Tag, error) {
	return a.repository.GetTagByID(ctx, tagID)
}

func (a *TagManagerProxyAdapter) CreateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	err := a.repository.CreateTag(ctx, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (a *TagManagerProxyAdapter) UpdateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	err := a.repository.UpdateTag(ctx, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (a *TagManagerProxyAdapter) DeleteTag(ctx context.Context, tagID string) error {
	return a.repository.DeleteTag(ctx, tagID)
}

/* relationships with memories */
func (a *TagManagerProxyAdapter) AssociateMemoryWithTag(ctx context.Context, memoryID string, tagID string) error {
	return a.repository.AssociateMemoryWithTag(ctx, memoryID, tagID)
}

func (a *TagManagerProxyAdapter) DisassociateMemoryFromTag(ctx context.Context, memoryID string, tagID string) error {
	return a.repository.DisassociateMemoryFromTag(ctx, memoryID, tagID)
}

func (a *TagManagerProxyAdapter) GetTagsByMemoryID(ctx context.Context, memoryID string) ([]models.Tag, error) {
	return a.repository.GetTagsByMemoryID(ctx, memoryID)
}

func (a *TagManagerProxyAdapter) GetMemoriesByTagID(ctx context.Context, tagID string) ([]string, error) {
	return a.repository.GetMemoriesByTagID(ctx, tagID)
}
