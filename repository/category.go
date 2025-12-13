package repository

import (
	"context"
	"fmt"
	"inventory-cli/model"
	"time"

	"github.com/jackc/pgx/v5"
)

// kontrak untuk data access
type CategoryRepository interface {
	Create(ctg *model.CategoryModel) error
	FindAll() ([]*model.CategoryModel, error)
	FindById(id int) (*model.CategoryModel, error)
	Update(ctg *model.CategoryModel) error
	Delete(id int) error
}

type categoryRepository struct {
	DB *pgx.Conn
}

func NewCategoryRepository(db *pgx.Conn) CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

func (repo *categoryRepository) Create(ctg *model.CategoryModel) error {
	query := `
	    INSERT INTO categories (name, description, created_at, updated_at)
    	VALUES ($1, $2, $3, $4)
    	RETURNING id, created_at, updated_at;
	`

	now := time.Now()

	err := repo.DB.QueryRow(
		context.Background(),
		query,
		ctg.Name,
		ctg.Description,
		now,
		now,
	).Scan(&ctg.ID, &ctg.CreatedAt, &ctg.UpdatedAt)

	if err != nil {
		return fmt.Errorf("repository: create category failed: %w", err)
	}

	return nil
}

func (repo *categoryRepository) FindAll() ([]*model.CategoryModel, error) {
	query := `
		SELECT id, name, description, created_at, updated_at
		FROM categories
		ORDER BY id ASC
	`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("repository: query categories failed: %w", err)
	}
	defer rows.Close()

	var categories []*model.CategoryModel

	for rows.Next() {
		var c model.CategoryModel
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Description,
			&c.CreatedAt,
			&c.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("repository: scan category failed: %w", err)
		}

		categories = append(categories, &c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository: rows iteration failed: %w", err)
	}

	return categories, nil
}

func (repo *categoryRepository) FindById(id int) (*model.CategoryModel, error) {
	query := `
		SELECT id, name, description, created_at, updated_at
		FROM categories
		WHERE id = $1;
	`

	var ctg model.CategoryModel

	err := repo.DB.QueryRow(context.Background(), query, id).
		Scan(
			&ctg.ID,
			&ctg.Name,
			&ctg.Description,
			&ctg.CreatedAt,
			&ctg.UpdatedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("repository: find category failed: %w", err)
	}

	return &ctg, nil
}

func (repo *categoryRepository) Update(ctg *model.CategoryModel) error {
	query := `
		UPDATE categories 
		SET name = $1, description = $2, updated_at = $3
		WHERE id = $4
		RETURNING updated_at;
	`

	now := time.Now()

	err := repo.DB.QueryRow(
		context.Background(),
		query,
		ctg.Name,
		ctg.Description,
		now,
		ctg.ID,
	).Scan(&ctg.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("repository: category id %d not found", ctg.ID)
		}
		return fmt.Errorf("repository: update category failed: %w", err)
	}

	return nil
}

func (repo *categoryRepository) Delete(id int) error {
	query := `
		DELETE FROM categories 
		WHERE id = $1;
	`

	cmdTag, err := repo.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("repository: delete category failed: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("repository: category id %d not found", id)
	}

	return nil
}
