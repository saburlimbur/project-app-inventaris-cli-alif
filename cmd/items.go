package cmd

import (
	"context"
	"fmt"
	"inventory-cli/database"
	"inventory-cli/dto"
	"inventory-cli/handler"
	"inventory-cli/model"
	"inventory-cli/repository"
	"inventory-cli/service"
	"inventory-cli/utils"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
)

type ItemsHandlerSetup struct {
	DB      *pgx.Conn
	Handler *handler.ItemsHandler
}

func setupItemsHandler() (*ItemsHandlerSetup, error) {
	db, err := database.InitDB()
	if err != nil {
		fmt.Printf("Gagal koneksi ke db")
		return nil, err
	}

	repo := repository.NewItemsRepository(db)
	svc := service.NewServiceItems(repo)
	h := handler.NewItemsHandler(svc)

	return &ItemsHandlerSetup{
		DB:      db,
		Handler: h,
	}, nil
}

// clean up helper
func (s *ItemsHandlerSetup) Close() {
	if s.DB != nil {
		s.DB.Close(context.Background())
	}
}

var listItems = &cobra.Command{
	Use:   "list-items",
	Short: "Menampilkan semua barang",
	Long:  "Mengambil dan menampilkan semua data barang dari database",
	Run: func(cmd *cobra.Command, args []string) {

		setup, err := setupItemsHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.ListsItems()
	},
}

var addItemCmd = &cobra.Command{
	Use:   "add-item",
	Short: "Tambahkan item/barang baru",
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")
		categoryID, _ := cmd.Flags().GetInt("category")
		price, _ := cmd.Flags().GetFloat64("price")
		dateStr, _ := cmd.Flags().GetString("date")
		usage, _ := cmd.Flags().GetInt("usage")

		// 2006-01-02 itu pattern time
		purchaseDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("Incorrect date format. Use YYYY-MM-DD")
			return
		}

		setup, err := setupItemsHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.CreateItem(&dto.CreateItemRequest{
			Name:         name,
			CategoryID:   categoryID,
			Price:        price,
			PurchaseDate: purchaseDate,
			UsageDays:    usage,
		})
	},
}

var detailItemyCmd = &cobra.Command{
	Use:   "detail-item",
	Short: "Melihat detail item berdasarkan ID",
	Run: func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetInt("id")

		if id == 0 {
			utils.PrintError("")
			fmt.Println("\nContoh: detail-item --id 1")
			return
		}

		setup, err := setupItemsHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.DetailItem(id)
	},
}

var deleteItemCmd = &cobra.Command{
	Use:   "delete-item",
	Short: "Hapus item berdasarkan ID",
	Long:  "Menghapus data item dari database berdasarkan ID yang diberikan",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")

		if id == 0 {
			utils.PrintError("Id item harus di isi")
			fmt.Println("\nContoh: delete-item --id=1")
			return
		}

		setup, err := setupItemsHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.DeleteItem(id)
	},
}

var searchItemCmd = &cobra.Command{
	Use:   "search-item",
	Short: "Cari item berdasarkan nama",
	Run: func(cmd *cobra.Command, args []string) {

		keyword, _ := cmd.Flags().GetString("name")

		if keyword == "" {
			fmt.Println("Keyword wajib diisi")
			fmt.Println("Contoh: search-item --name laptop")
			return
		}

		setup, err := setupItemsHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.SearchItems(keyword)
	},
}

var updateItemCmd = &cobra.Command{
	Use:   "update-item",
	Short: "Update data barang",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		price, _ := cmd.Flags().GetFloat64("price")
		categoryID, _ := cmd.Flags().GetInt("category")
		usage, _ := cmd.Flags().GetInt("usage")
		date, _ := cmd.Flags().GetString("date")

		pd, _ := time.Parse("2006-01-02", date)

		setup, _ := setupItemsHandler()
		defer setup.Close()

		setup.Handler.UpdateItem(&model.ItemsModel{
			ID:           id,
			Name:         name,
			Price:        price,
			CategoryID:   categoryID,
			PurchaseDate: pd,
			UsageDays:    usage,
		})
	},
}

// barang yang perlu diganti lebih dari 100 hari
var needReplacementCmd = &cobra.Command{
	Use:   "need-replacement",
	Short: "Menampilkan barang yang perlu diganti",
	Long:  "Menampilkan daftar barang dengan usage_days > 100",
	Run: func(cmd *cobra.Command, args []string) {

		setup, err := setupItemsHandler()
		if err != nil {
			fmt.Println("Gagal setup handler")
			return
		}
		defer setup.Close()

		setup.Handler.ItemsNeedReplacement()
	},
}

var investmentSummaryCmd = &cobra.Command{
	Use:   "investment-summary",
	Short: "Laporan total investasi dan depresiasi",
	Run: func(cmd *cobra.Command, args []string) {

		setup, err := setupItemsHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.InvestmentSummary()
	},
}

var investmentDetailCmd = &cobra.Command{
	Use:   "investment-detail",
	Short: "Detail investasi barang berdasarkan ID",
	Run: func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetInt("id")
		if id == 0 {
			fmt.Println("Contoh: investment-detail --id 1")
			return
		}

		setup, err := setupItemsHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.InvestmentDetail(id)
	},
}

func init() {
	rootCmd.AddCommand(listItems)
	rootCmd.AddCommand(addItemCmd)
	rootCmd.AddCommand(detailItemyCmd)
	rootCmd.AddCommand(deleteItemCmd)
	rootCmd.AddCommand(searchItemCmd)
	rootCmd.AddCommand(updateItemCmd)
	rootCmd.AddCommand(needReplacementCmd)

	rootCmd.AddCommand(investmentSummaryCmd)
	rootCmd.AddCommand(investmentDetailCmd)

	// add item
	addItemCmd.Flags().IntP("category", "c", 0, "Category ID (required)")
	addItemCmd.Flags().StringP("name", "n", "", "Nama item (required)")
	addItemCmd.Flags().Float64P("price", "p", 0, "Harga item")
	addItemCmd.Flags().StringP("date", "d", "", "Tanggal beli (YYYY-MM-DD)")
	addItemCmd.Flags().IntP("usage", "u", 0, "Jumlah hari pemakaian")

	addItemCmd.MarkFlagRequired("category")
	addItemCmd.MarkFlagRequired("name")
	addItemCmd.MarkFlagRequired("price")
	addItemCmd.MarkFlagRequired("date")

	// detail
	detailItemyCmd.Flags().IntP("id", "i", 0, "ID item (required)")
	detailItemyCmd.MarkFlagRequired("id")

	// delete
	deleteItemCmd.Flags().IntP("id", "i", 0, "ID item (required)")
	deleteItemCmd.MarkFlagRequired("id")

	// search item/barang
	searchItemCmd.Flags().StringP("name", "n", "", "Nama item yang dicari")

	// update-item
	updateItemCmd.Flags().IntP("id", "i", 0, "ID item (required)")
	updateItemCmd.Flags().IntP("category", "c", 0, "Category ID")
	updateItemCmd.Flags().StringP("name", "n", "", "Nama item")
	updateItemCmd.Flags().Float64P("price", "p", 0, "Harga item")
	updateItemCmd.Flags().StringP("date", "d", "", "Tanggal beli (YYYY-MM-DD)")
	updateItemCmd.Flags().IntP("usage", "u", 0, "Jumlah hari pemakaian")

	// invesment
	investmentDetailCmd.Flags().IntP("id", "i", 0, "ID item (required)")
	investmentDetailCmd.MarkFlagRequired("id")
}
