package repositories

import (
	"server/internal/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CategoryRepo interface {
	Create(ctx context.Context, category *models.Category) error
	GetByID(ctx context.Context, id int64) (*models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int64) error
}

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) Create(ctx context.Context, category *models.Category) error {
	query := `
		INSERT INTO categories (name, description, is_active)
		VALUES (:name, :description, :is_active)
		RETURNING category_id
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()
	
	return stmt.GetContext(ctx, category, category)
}

func (r *categoryRepo) GetByID(ctx context.Context, id int64) (*models.Category, error) {
	var category models.Category
	query := `SELECT * FROM categories WHERE category_id = $1`
	err := r.db.GetContext(ctx, &category, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &category, err
}

func (r *categoryRepo) GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	query := `SELECT * FROM categories ORDER BY name`
	err := r.db.SelectContext(ctx, &categories, query)
	return categories, err
}

func (r *categoryRepo) Update(ctx context.Context, category *models.Category) error {
	query := `
		UPDATE categories 
		SET name = :name, 
		    description = :description,
		    is_active = :is_active
		WHERE category_id = :category_id
	`
	result, err := r.db.NamedExecContext(ctx, query, category)
	if err != nil {
		return err
	}
	
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *categoryRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM categories WHERE category_id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}