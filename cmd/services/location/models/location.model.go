package models

type LocationResponse struct {
	ID         string  `json:"id"`
	LocationID string  `json:"location_id"`
	Location   string  `json:"location"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
}

type CreateLocationRequest struct {
	Location  string  `json:"location" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	City      string  `json:"city" binding:"required"`
	Country   string  `json:"country" binding:"required"`
}

type UpdateLocationRequest struct {
	Location  string   `json:"location,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	City      string   `json:"city,omitempty"`
	Country   string   `json:"country,omitempty"`
}
