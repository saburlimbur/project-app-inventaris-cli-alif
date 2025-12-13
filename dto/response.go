package dto

import "time"

type CategoryResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemResponse struct {
	ID           int       `json:"id"`
	CategoryID   int       `json:"category_id"`
	CategoryName string    `json:"category_name,omitempty"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
	UsageDays    int       `json:"usage_days"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
