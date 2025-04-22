package models

type TagResponse struct {
	ID    string `json:"id"`
	TagID string `json:"tag_id"`
	Name  string `json:"name"`
}

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTagRequest struct {
	Name string `json:"name" binding:"required"`
}