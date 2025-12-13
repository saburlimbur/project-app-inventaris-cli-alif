package utils

import (
	"fmt"
	"inventory-cli/model"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/olekukonko/tablewriter"
)

func PrintCategoryTable(categories []*model.CategoryModel) {
	if len(categories) == 0 {
		fmt.Println("No categories found.")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("CATEGORY LIST")

	// Header tanpa ANSI
	t.AppendHeader(table.Row{"ID", "NAME", "DESCRIPTION", "CREATED AT"})

	for _, c := range categories {
		t.AppendRow(table.Row{
			c.ID,
			c.Name,
			c.Description,
			c.CreatedAt,
		})
	}

	t.SetStyle(table.StyleLight)
	t.Render()
}

func PrintItemsTable(items []*model.ItemsModel) {
	if len(items) == 0 {
		fmt.Println("No items found.")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("ITEMS LIST")

	t.AppendHeader(table.Row{"ID", "CATEGORY ID", "NAME", "PRICE", "PURCHASE DATE", "USAGE DAYS", "CREATED AT"})

	for _, itm := range items {
		t.AppendRow(table.Row{
			itm.ID,
			itm.CategoryID,
			itm.Name,
			fmt.Sprintf("%.0f", itm.Price),
			itm.PurchaseDate,
			itm.UsageDays,
			itm.CreatedAt,
		})
	}

	t.SetStyle(table.StyleLight)
	t.Render()
}

func PrintCategoryDetail(categories []*model.CategoryModel) {
	fmt.Println("DETAIL KATEGORI:")

	table := tablewriter.NewWriter(os.Stdout)

	table.Append([]string{"ID", "Name", "Description", "Created At"})

	for _, c := range categories {
		row := []string{
			fmt.Sprintf("%d", c.ID),
			c.Name,
			c.Description,
			c.CreatedAt.String(),
		}

		table.Append(row)
	}

	table.Render()

	fmt.Println("============================================================")
}
