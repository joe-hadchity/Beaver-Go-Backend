package services

import (
	"server/internal/models"
	"server/internal/repositories"
	"context"
	"errors"
	"fmt"
)

type CategoryService struct {
	repo repositories.CategoryRepo
}

func NewCategoryService(repo repositories.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *models.Category) (*models.Category, error) {
	if req.Name == "" {
		return nil, errors.New("category name cannot be empty")
	}

	existing, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to verify category: %w", err)
	}

	for _, c := range existing {
		if c.Name == req.Name {
			return nil, errors.New("category name already exists")
		}
	}

	if err := s.repo.Create(ctx, req); err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return req, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, id int64) (*models.Category, error) {
	category, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}
	return category, nil
}

func (s *CategoryService) ListCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}
	return categories, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, req *models.Category) error {
	if req.ID == 0 {
		return errors.New("invalid category ID")
	}
	
	if err := s.repo.Update(ctx, req); err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}
	return nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("invalid category ID")
	}
	
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil
}