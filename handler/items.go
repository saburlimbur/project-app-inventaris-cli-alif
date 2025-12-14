package handler

import (
	"fmt"
	"inventory-cli/model"
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

func (h *ItemsHandler) DetailItem(id int) {
	itm, err := h.ItemsSvc.GetItemByID(id)
	if err != nil {
		fmt.Println("Error getting item detail:", err)
		return
	}

	if itm == nil {
		fmt.Println("item not found.")
		return
	}

	utils.PrintItemDetailTable(itm)
}

func (h *ItemsHandler) SearchItems(keyword string) {
	items, err := h.ItemsSvc.SearchItems(keyword)
	if err != nil {
		fmt.Println("Error searching items:", err)
		return
	}

	if len(items) == 0 {
		fmt.Println("Item not found.")
		return
	}

	utils.PrintItemsTable(items)
}

func (h *ItemsHandler) UpdateItem(itm *model.ItemsModel) {
	err := h.ItemsSvc.UpdateItem(itm)
	if err != nil {
		fmt.Println("Gagal update item:", err)
		return
	}
	fmt.Println("Item berhasil diperbarui")
}

func (h *ItemsHandler) ItemsNeedReplacement() {
	items, err := h.ItemsSvc.GetItemsNeedReplacement()
	if err != nil {
		fmt.Println("Gagal mengambil data:", err)
		return
	}

	if len(items) == 0 {
		fmt.Println("Tidak ada barang yang perlu diganti")
		return
	}

	utils.PrintItemsTable(items)
}

func (h *ItemsHandler) InvestmentSummary() {
	res, err := h.ItemsSvc.GetInvestmentSummary()
	if err != nil {
		fmt.Println("Gagal menghitung investasi:", err)
		return
	}

	utils.PrintInvestmentSummaryTable(res)
}

func (h *ItemsHandler) InvestmentDetail(id int) {
	res, err := h.ItemsSvc.GetInvestmentDetail(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.PrintInvestmentDetailTable(res)
}
