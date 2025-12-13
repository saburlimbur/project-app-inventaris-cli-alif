package handler

import (
	"fmt"
	"inventory-cli/service"
	"inventory-cli/utils"
)

type ItemsHandler struct {
	ItemsSvc service.ItemsService
}

func NewItemsHandler(svc service.ItemsService) *ItemsHandler {
	return &ItemsHandler{
		ItemsSvc: svc,
	}
}

func (h *ItemsHandler) ListsItems() {
	items, err := h.ItemsSvc.Lists()

	if err != nil {
		fmt.Printf("Failed getting item lists: %v\n", err)
	}

	if len(items) == 0 {
		fmt.Println("Items not found")
		return
	}

	utils.PrintItemsTable(items)
}

func (h *ItemsHandler) CreateItem() {
	
}
