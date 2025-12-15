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
	FindById(id int) (*model.ItemsModel, error)
	Delete(id int) error
	SearchByName(keyword string) ([]*model.ItemsModel, error)
	Update(itm *model.ItemsModel) error
	FindNeedReplacement() ([]*model.ItemsModel, error)
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
	RETURNING id, created_at, updated_at;
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

func (repo *itemsRepository) FindById(id int) (*model.ItemsModel, error) {
	query := `
		SELECT id, category_id, name, price, purchase_date, usage_days, created_at, updated_at
		FROM items
		WHERE id = $1;
	`

	var itm model.ItemsModel

	err := repo.DB.QueryRow(context.Background(), query, id).
		Scan(
			&itm.ID,
			&itm.CategoryID,
			&itm.Name,
			&itm.Price,
			&itm.PurchaseDate,
			&itm.UsageDays,
			&itm.CreatedAt,
			&itm.UpdatedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("repository: find item failed: %w", err)
	}

	return &itm, nil
}

func (repo *itemsRepository) SearchByName(keyword string) ([]*model.ItemsModel, error) {
	query := `
		SELECT id, category_id, name, price, purchase_date, usage_days, created_at, updated_at
		FROM items
		WHERE name ILIKE '%' || $1 || '%'
		ORDER BY id ASC;
	`

	rows, err := repo.DB.Query(context.Background(), query, keyword)
	if err != nil {
		return nil, fmt.Errorf("repository: search item failed: %w", err)
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

	return items, nil
}

func (repo *itemsRepository) Update(itm *model.ItemsModel) error {
	query := `
		UPDATE items
		SET category_id = $1,
		    name = $2,
		    price = $3,
		    purchase_date = $4,
		    usage_days = $5,
		    updated_at = NOW()
		WHERE id = $6
	`

	ct, err := repo.DB.Exec(
		context.Background(),
		query,
		itm.CategoryID,
		itm.Name,
		itm.Price,
		itm.PurchaseDate,
		itm.UsageDays,
		itm.ID,
	)

	if err != nil {
		return fmt.Errorf("repository: update item failed: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// barang atau item yang usage_days > 100
func (repo *itemsRepository) FindNeedReplacement() ([]*model.ItemsModel, error) {
	query := `
		SELECT 
			id,
			category_id,
			name,
			price,
			purchase_date,
			usage_days,
			created_at,
			updated_at
		FROM items
		WHERE usage_days > 100
		ORDER BY usage_days DESC
	`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("repository: query need replacement failed: %w", err)
	}
	defer rows.Close()

	var items []*model.ItemsModel

	for rows.Next() {
		var itm model.ItemsModel
		err := rows.Scan(
			&itm.ID,
			&itm.CategoryID,
			&itm.Name,
			&itm.Price,
			&itm.PurchaseDate,
			&itm.UsageDays,
			&itm.CreatedAt,
			&itm.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("repository: scan failed: %w", err)
		}
		items = append(items, &itm)
	}

	return items, nil
}

func (repo *itemsRepository) Delete(id int) error {
	query := `
	DELETE FROM items
	WHERE id = $1;
	`

	rows, err := repo.DB.Exec(context.Background(), query, id)

	if err != nil {
		return fmt.Errorf("repository: delete item failed: %w", err)
	}

	if rows.RowsAffected() == 0 {
		return fmt.Errorf("repository: item id %d not found", id)
	}

	return nil
}
