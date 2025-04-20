package models

import (
	"recorderis/internals/utils"
	"time"
)

/* API Responses */
type MemoryResponse struct {
	ID           string                `json:"id"`
	MemoryID     string                `json:"memory_id"`
	Title        string                `json:"title"`
	Date         time.Time             `json:"date"`
	IsPublic     bool                  `json:"is_public"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	Descriptions []DescriptionResponse `json:"descriptions,omitempty"`
	Photos       []PhotoResponse       `json:"photos,omitempty"`
}

type DescriptionResponse struct {
	ID            string `json:"id"`
	DescriptionID string `json:"description_id"`
	Text          string `json:"text"`
	Index         int    `json:"index"`
	Version       int    `json:"version"`
}

type PhotoResponse struct {
	ID       string `json:"id"`
	PhotoID  string `json:"photo_id"`
	FileName string `json:"file_name"`
	URL      string `json:"url"`
	Index    int    `json:"index"`
}

/* API Requests */
type CreateMemoryRequest struct {
	Title        string                     `json:"title" binding:"required"`
	Date         utils.JSONTime             `json:"date" binding:"required"`
	IsPublic     bool                       `json:"is_public"`
	Descriptions []CreateDescriptionRequest `json:"descriptions,omitempty"`
}

type UpdateMemoryRequest struct {
	Title    string         `json:"title,omitempty"`
	Date     utils.JSONTime `json:"date,omitempty"`
	IsPublic *bool          `json:"is_public,omitempty"`
	Index    *int           `json:"index,omitempty"`
}

type CreateDescriptionRequest struct {
	Text  string `json:"text" binding:"required"`
	Index int    `json:"index"`
}

type UpdateDescriptionRequest struct {
	Text  string `json:"text,omitempty"`
	Index int    `json:"index,omitempty"`
}

type CreatePhotoRequest struct {
	FileName string `json:"file_name" binding:"required"`
	URL      string `json:"url" binding:"required"`
	Index    int    `json:"index"`
}
