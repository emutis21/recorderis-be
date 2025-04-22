package drivens

import (
	"context"
	repo_adapters "recorderis/cmd/services/repository/adapters/drivers"
	repo_models "recorderis/cmd/services/repository/models"
	"recorderis/cmd/services/tags/models"
	"recorderis/cmd/services/tags/ports/drivens"
	"strconv"
)

var _ drivens.ForTagRepository = (*TagRepositoryAdapter)(nil)

type TagRepositoryAdapter struct {
	tagRepo *repo_adapters.TagManagerProxyAdapter
}

func NewTagRepositoryAdapter(tagRepo *repo_adapters.TagManagerProxyAdapter) *TagRepositoryAdapter {
	return &TagRepositoryAdapter{
		tagRepo: tagRepo,
	}
}

func (a *TagRepositoryAdapter) GetTags(ctx context.Context) ([]models.TagResponse, error) {
	tags, err := a.tagRepo.GetTags(ctx)
	if err != nil {
		return nil, err
	}

	var response []models.TagResponse
	for _, tag := range tags {
		response = append(response, mapTagToResponse(&tag))
	}

	return response, nil
}

func (a *TagRepositoryAdapter) GetTagByID(ctx context.Context, tagID string) (*models.TagResponse, error) {
	tag, err := a.tagRepo.GetTagByID(ctx, tagID)
	if err != nil {
		return nil, err
	}

	response := mapTagToResponse(tag)
	return &response, nil
}

func (a *TagRepositoryAdapter) CreateTag(ctx context.Context, req *models.CreateTagRequest) (*models.TagResponse, error) {
	tag := &repo_models.Tag{
		Name: req.Name,
	}

	createdTag, err := a.tagRepo.CreateTag(ctx, tag)
	if err != nil {
		return nil, err
	}

	response := mapTagToResponse(createdTag)
	return &response, nil
}

func (a *TagRepositoryAdapter) UpdateTag(ctx context.Context, tagID string, req *models.UpdateTagRequest) (*models.TagResponse, error) {
	tag, err := a.tagRepo.GetTagByID(ctx, tagID)
	if err != nil {
		return nil, err
	}

	tag.Name = req.Name

	updatedTag, err := a.tagRepo.UpdateTag(ctx, tag)
	if err != nil {
		return nil, err
	}

	response := mapTagToResponse(updatedTag)
	return &response, nil
}

func (a *TagRepositoryAdapter) DeleteTag(ctx context.Context, tagID string) error {
	return a.tagRepo.DeleteTag(ctx, tagID)
}

func (a *TagRepositoryAdapter) AssociateMemoryWithTag(ctx context.Context, memoryID string, tagID string) error {
	return a.tagRepo.AssociateMemoryWithTag(ctx, memoryID, tagID)
}

func (a *TagRepositoryAdapter) DisassociateMemoryFromTag(ctx context.Context, memoryID string, tagID string) error {
	return a.tagRepo.DisassociateMemoryFromTag(ctx, memoryID, tagID)
}

func (a *TagRepositoryAdapter) GetTagsByMemoryID(ctx context.Context, memoryID string) ([]models.TagResponse, error) {
	tags, err := a.tagRepo.GetTagsByMemoryID(ctx, memoryID)
	if err != nil {
		return nil, err
	}

	var response []models.TagResponse
	for _, tag := range tags {
		response = append(response, mapTagToResponse(&tag))
	}

	return response, nil
}

func (a *TagRepositoryAdapter) GetMemoriesByTagID(ctx context.Context, tagID string) ([]string, error) {
	return a.tagRepo.GetMemoriesByTagID(ctx, tagID)
}

func mapTagToResponse(tag *repo_models.Tag) models.TagResponse {
	return models.TagResponse{
		ID:    strconv.Itoa(tag.ID),
		TagID: tag.TagID,
		Name:  tag.Name,
	}
}
