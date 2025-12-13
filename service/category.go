package service

import (
	"errors"
	"fmt"
	"inventory-cli/model"
	"inventory-cli/repository"
	"inventory-cli/utils"
	"strings"
)

type CategoryService interface {
	CreateCategory(ctg *model.CategoryModel) error
	GetAllCategories() ([]*model.CategoryModel, error)
	GetCategoryByID(id int) (*model.CategoryModel, error)
	UpdateCategory(ctg *model.CategoryModel) error
	DeleteCategory(id int) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) validateCategory(ctg *model.CategoryModel) error {
	if ctg == nil {
		return errors.New("category cannot be nil")
	}

	name := strings.TrimSpace(ctg.Name)
	if name == "" {
		return utils.ErrCategoryNameEmpty
	}

	if len(name) > 100 {
		return utils.ErrCategoryNameTooLong
	}

	if len(ctg.Description) > 500 {
		return utils.ErrCategoryDescTooLong
	}

	return nil
}

func (s *categoryService) CreateCategory(ctg *model.CategoryModel) error {
	if err := s.validateCategory(ctg); err != nil {
		return err
	}

	ctg.Name = strings.TrimSpace(ctg.Name)
	ctg.Description = strings.TrimSpace(ctg.Description)

	// instance ke repo
	if err := s.repo.Create(ctg); err != nil {
		return fmt.Errorf("service: failed to create category: %w", err)
	}

	return nil
}

func (s *categoryService) GetAllCategories() ([]*model.CategoryModel, error) {
	categories, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("service: failed to get categories: %w", err)
	}

	return categories, nil
}

func (s *categoryService) GetCategoryByID(id int) (*model.CategoryModel, error) {
	if id <= 0 {
		return nil, utils.ErrInvalidCategoryID
	}

	ctg, err := s.repo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get category: %w", err)
	}

	if ctg == nil {
		return nil, utils.ErrCategoryNotFound
	}

	return ctg, nil
}

func (s *categoryService) UpdateCategory(ctg *model.CategoryModel) error {
	// Validasi business rules
	if err := s.validateCategory(ctg); err != nil {
		return err
	}

	// Cek existence
	existing, err := s.repo.FindById(ctg.ID)
	if err != nil {
		return fmt.Errorf("service: failed to check category: %w", err)
	}

	if existing == nil {
		return utils.ErrCategoryNotFound
	}

	// Normalize
	ctg.Name = strings.TrimSpace(ctg.Name)
	ctg.Description = strings.TrimSpace(ctg.Description)

	// Update
	if err := s.repo.Update(ctg); err != nil {
		return fmt.Errorf("service: failed to update category: %w", err)
	}

	return nil
}

func (s *categoryService) DeleteCategory(id int) error {
	if id <= 0 {
		return utils.ErrInvalidCategoryID
	}

	// Cek existence
	existing, err := s.repo.FindById(id)
	if err != nil {
		return fmt.Errorf("service: failed to check category: %w", err)
	}

	if existing == nil {
		return utils.ErrCategoryNotFound
	}

	// Delete
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("service: failed to delete category: %w", err)
	}

	return nil
}
