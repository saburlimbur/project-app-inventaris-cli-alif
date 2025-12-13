package model

import "time"

type ItemsModel struct {
	Model
	ID           int       `json:"id"`
	CategoryID   int       `json:"category_id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
	UsageDays    int       `json:"usage_days"`
}
