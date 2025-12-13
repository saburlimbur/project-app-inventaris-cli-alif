package cmd

import (
	"context"
	"fmt"
	"inventory-cli/database"
	"inventory-cli/handler"
	"inventory-cli/repository"
	"inventory-cli/service"
	"inventory-cli/utils"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
)

type CategoryHandlerSetup struct {
	DB      *pgx.Conn
	Handler *handler.CategoryHandler
}

func setupCategoryHandler() (*CategoryHandlerSetup, error) {
	db, err := database.InitDB()
	if err != nil {
		fmt.Printf("Gagal koneksi ke database")
		return nil, err
	}

	// init layer inject dependensi
	repo := repository.NewCategoryRepository(db)
	svc := service.NewCategoryService(repo)
	h := handler.NewCategoryHandler(svc)

	return &CategoryHandlerSetup{
		DB:      db,
		Handler: h,
	}, nil
}

// clean up helper
func (s *CategoryHandlerSetup) Close() {
	if s.DB != nil {
		s.DB.Close(context.Background())
	}
}

var addCategoryCmd = &cobra.Command{
	Use:   "add-category",
	Short: "Tambahkan kategori baru",
	Long:  "Menambahkan data kategori baru ke database inventory",
	Example: `  inventory-cli add-category --name "Elektronik" --desc "Barang elektronik"
  inventory-cli add-category -n "Makanan" -d "Produk makanan"`,
	Run: func(cmd *cobra.Command, args []string) {
		//  flags locally (tidak pakai global var)
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("desc")

		// setup dependencies
		setup, err := setupCategoryHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		// execute handler
		setup.Handler.CreateCategory(name, desc)
	},
}

var listCategoryCmd = &cobra.Command{
	Use:     "list-category",
	Short:   "Menampilkan semua kategori",
	Long:    "Mengambil dan menampilkan semua data kategori dari database",
	Example: `  inventory-cli list-category`,
	Run: func(cmd *cobra.Command, args []string) {
		setup, err := setupCategoryHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.ListsCategory()
	},
}

var detailCategoryCmd = &cobra.Command{
	Use:   "detail-category",
	Short: "Melihat detail kategori berdasarkan ID",
	Run: func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetInt("id")

		if id == 0 {
			utils.PrintError("")
			fmt.Println("\nContoh: detail-category --id 1")
			return
		}

		setup, err := setupCategoryHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.DetailCategory(id)
	},
}

var updateCategoryCmd = &cobra.Command{
	Use: "update-category",
	Example: `  inventory-cli update-category --id 1 --name "Elektronik Baru" --desc "Deskripsi baru"
  inventory-cli update-category -i 2 -n "Makanan" -d "Produk makanan dan minuman"`,
	Run: func(cmd *cobra.Command, args []string) {
		//  flags locally
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("desc")

		if id == 0 {
			utils.PrintError("ID kategori harus diisi")
			fmt.Println("\nContoh: update-category --id 1 --name 'Nama' --desc 'Deskripsi'")
			return
		}

		setup, err := setupCategoryHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.UpdateCategory(id, name, desc)
	},
}

var deleteCategoryCmd = &cobra.Command{
	Use:   "delete-category",
	Short: "Hapus kategori berdasarkan ID",
	Long:  "Menghapus data kategori dari database berdasarkan ID yang diberikan",
	Example: `  inventory-cli delete-category --id 1
  inventory-cli delete-category -i 5`,
	Run: func(cmd *cobra.Command, args []string) {
		// flag locally
		id, _ := cmd.Flags().GetInt("id")

		// validasi input
		if id == 0 {
			utils.PrintError("ID kategori harus diisi")
			fmt.Println("\nContoh: delete-category --id 1")
			return
		}

		setup, err := setupCategoryHandler()
		if err != nil {
			return
		}
		defer setup.Close()

		setup.Handler.DeleteCategory(id)
	},
}

func init() {
	rootCmd.AddCommand(addCategoryCmd)
	rootCmd.AddCommand(listCategoryCmd)
	rootCmd.AddCommand(detailCategoryCmd)
	rootCmd.AddCommand(updateCategoryCmd)
	rootCmd.AddCommand(deleteCategoryCmd)

	// add
	addCategoryCmd.Flags().StringP("name", "n", "", "Nama kategori (required)")
	addCategoryCmd.Flags().StringP("desc", "d", "", "Deskripsi kategori (required)")
	addCategoryCmd.MarkFlagRequired("name")
	addCategoryCmd.MarkFlagRequired("desc")

	// detail id
	detailCategoryCmd.Flags().IntP("id", "i", 0, "ID kategori (required)")
	detailCategoryCmd.MarkFlagRequired("id")

	// update id
	updateCategoryCmd.Flags().IntP("id", "i", 0, "ID kategori (required)")
	updateCategoryCmd.Flags().StringP("name", "n", "", "Nama kategori baru (required)")
	updateCategoryCmd.Flags().StringP("desc", "d", "", "Deskripsi kategori baru (required)")
	updateCategoryCmd.MarkFlagRequired("id")
	updateCategoryCmd.MarkFlagRequired("name")
	updateCategoryCmd.MarkFlagRequired("desc")

	// delete id
	deleteCategoryCmd.Flags().IntP("id", "i", 0, "ID kategori (required)")
	deleteCategoryCmd.MarkFlagRequired("id")
}
