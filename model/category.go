package model

type CategoryModel struct {
	Model
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
