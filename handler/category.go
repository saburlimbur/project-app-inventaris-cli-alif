package handler

import (
	"fmt"
	"inventory-cli/model"
	"inventory-cli/service"
	"inventory-cli/utils"
)

type CategoryHandler struct {
	CategorySvc service.CategoryService
}

func NewCategoryHandler(svc service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategorySvc: svc,
	}
}

func (h *CategoryHandler) CreateCategory(name, desc string) error {
	ctg := &model.CategoryModel{
		Name:        name,
		Description: desc,
	}

	if err := h.CategorySvc.CreateCategory(ctg); err != nil {
		fmt.Println("Failed to create category:", err)
		return err
	}

	fmt.Println("Category created successfully.")
	return nil
}

func (h *CategoryHandler) ListsCategory() {
	categories, err := h.CategorySvc.GetAllCategories()
	if err != nil {
		fmt.Printf("Failed getting category lists: %v\n", err)
		return
	}

	if len(categories) == 0 {
		fmt.Println("Category not found")
		return
	}

	utils.PrintCategoryTable(categories)
}

func (h *CategoryHandler) DetailCategory(id int) {
	ctg, err := h.CategorySvc.GetCategoryByID(id)
	if err != nil {
		fmt.Println("Error getting category detail:", err)
		return
	}

	if ctg == nil {
		fmt.Println("Category not found.")
		return
	}

	fmt.Println("===== DETAIL CATEGORY =====")
	fmt.Println("ID:", ctg.ID)
	fmt.Println("Name:", ctg.Name)
	fmt.Println("Description:", ctg.Description)
	fmt.Println("Created At:", ctg.CreatedAt.String())
	fmt.Println("Updated At:", ctg.UpdatedAt.String())
	fmt.Println("===========================")
}

func (h *CategoryHandler) UpdateCategory(id int, name, desc string) error {
	ctg := &model.CategoryModel{
		ID:          id,
		Name:        name,
		Description: desc,
	}

	if err := h.CategorySvc.UpdateCategory(ctg); err != nil {
		fmt.Println("Failed to update category:", err)
		return err
	}

	fmt.Println("Category updated successfully.")

	return nil
}

func (h *CategoryHandler) DeleteCategory(id int) error {
	if err := h.CategorySvc.DeleteCategory(id); err != nil {
		fmt.Println("Failed to delete category:", err)
		return err
	}

	fmt.Printf("Category with ID %d deleted successfully.\n", id)
	return nil
}
