package repository

import (
	"context"
	"fmt"
	"inventory-cli/model"

	"github.com/jackc/pgx/v5"
)

// kontrak untuk data access layer
type ItemsRepository interface {
	Create(itm *model.ItemsModel) error
	FindAll() ([]*model.ItemsModel, error)
}

type itemsRepository struct {
	// DB database.PgxIface
	DB *pgx.Conn
}

func NewItemsRepository(db *pgx.Conn) ItemsRepository {
	return &itemsRepository{
		DB: db,
	}
}

func (repo *itemsRepository) Create(itm *model.ItemsModel) error {
	query := `
		INSERT INTO items (category_id, name, price, purchase_date, usage_days)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, purchase_date, created_at, updated_at;
	`

	err := repo.DB.QueryRow(
		context.Background(),
		query,
		itm.CategoryID,
		itm.Name,
		itm.Price,
		itm.PurchaseDate,
		itm.UsageDays,
	).Scan(
		&itm.ID,
		&itm.PurchaseDate,
		&itm.CreatedAt,
		&itm.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("repository: create item failed: %w", err)
	}

	return nil
}

func (repo *itemsRepository) FindAll() ([]*model.ItemsModel, error) {
	query := `
		SELECT id, category_id, name, price, purchase_date, usage_days, created_at, updated_at
		FROM items
		ORDER BY id ASC
	`

	rows, err := repo.DB.Query(context.Background(), query)

	if err != nil {
		return nil, fmt.Errorf("repository: query item failed: %w", err)
	}

	defer rows.Close()

	var items []*model.ItemsModel

	for rows.Next() {
		var i model.ItemsModel
		err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Price,
			&i.PurchaseDate,
			&i.UsageDays,
			&i.CreatedAt,
			&i.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("repository: scan item failed: %w", err)
		}

		items = append(items, &i)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository: rows iteration failed: %w", err)
	}

	return items, nil
}
