package cmd

import (
	"context"
	"fmt"
	"inventory-cli/database"
	"inventory-cli/handler"
	"inventory-cli/repository"
	"inventory-cli/service"

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

func init() {
	rootCmd.AddCommand(listItems)
}
