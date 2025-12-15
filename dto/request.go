package dto

import "time"

type CreateItemRequest struct {
	CategoryID   int
	Name         string
	Price        float64
	PurchaseDate time.Time
	UsageDays    int
}
