package utils

import (
	"fmt"
	"inventory-cli/dto"
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

func PrintCategoryDetailTable(c *model.CategoryModel) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("CATEGORY DETAIL")

	t.AppendHeader(table.Row{"ID", "NAME", "DESCRIPTION", "CREATED AT", "UPDATED AT"})

	t.AppendRow(table.Row{
		c.ID,
		c.Name,
		c.Description,
		c.CreatedAt,
		c.UpdatedAt,
	})

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

func PrintItemDetailTable(itm *model.ItemsModel) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("ITEM DETAIL")

	t.AppendHeader(table.Row{
		"ID",
		"CATEGORY ID",
		"NAME",
		"PRICE",
		"PURCHASE DATE",
		"USAGE DAYS",
		"CREATED AT",
	})

	t.AppendRow(table.Row{
		itm.ID,
		itm.CategoryID,
		itm.Name,
		fmt.Sprintf("%.0f", itm.Price),
		itm.PurchaseDate.Format("2006-01-02"),
		itm.UsageDays,
		itm.CreatedAt.Format("2006-01-02 15:04:05"),
	})

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

func PrintInvestmentSummaryTable(res *dto.InvestmentSummaryResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("INVESTMENT SUMMARY")

	t.AppendHeader(table.Row{
		"TOTAL INITIAL VALUE",
		"TOTAL CURRENT VALUE",
		"TOTAL DEPRECIATION",
	})

	t.AppendRow(table.Row{
		fmt.Sprintf("Rp %.0f", res.TotalInitialValue),
		fmt.Sprintf("Rp %.0f", res.TotalCurrentValue),
		fmt.Sprintf("Rp %.0f",
			res.TotalInitialValue-res.TotalCurrentValue),
	})

	t.SetStyle(table.StyleLight)
	t.Render()
}

func PrintInvestmentDetailTable(res *dto.InvestmentDetailResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("INVESTMENT DETAIL")

	t.AppendHeader(table.Row{
		"ITEM ID",
		"ITEM NAME",
		"INITIAL VALUE",
		"CURRENT VALUE",
		"DEPRECIATION",
		"YEARS USED",
	})

	t.AppendRow(table.Row{
		res.ItemID,
		res.ItemName,
		fmt.Sprintf("Rp %.0f", res.InitialValue),
		fmt.Sprintf("Rp %.0f", res.CurrentValue),
		fmt.Sprintf("Rp %.0f", res.Depreciation),
		res.YearsUsed,
	})

	t.SetStyle(table.StyleLight)
	t.Render()
}
