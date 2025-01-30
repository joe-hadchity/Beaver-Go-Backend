package repositories

import (
	"server/internal/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ServiceRepo interface {
	Create(ctx context.Context, service *models.Service) error
	GetByID(ctx context.Context, id int64) (*models.Service, error)
	GetByCategory(ctx context.Context, categoryID int64) ([]models.Service, error)
	GetAll(ctx context.Context) ([]models.Service, error)
	Update(ctx context.Context, service *models.Service) error
	Delete(ctx context.Context, id int64) error
}

type serviceRepo struct {
	db *sqlx.DB
}

func NewServiceRepo(db *sqlx.DB) ServiceRepo {
	return &serviceRepo{db: db}
}

func (r *serviceRepo) Create(ctx context.Context, service *models.Service) error {
	query := `
		INSERT INTO services (category_id, name, description, is_active)
		VALUES (:category_id, :name, :description, :is_active)
		RETURNING service_id
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()
	
	return stmt.GetContext(ctx, service, service)
}

func (r *serviceRepo) GetByID(ctx context.Context, id int64) (*models.Service, error) {
	var service models.Service
	query := `
		SELECT s.*, c.name as category_name 
		FROM services s
		JOIN categories c ON s.category_id = c.category_id
		WHERE service_id = $1
	`
	err := r.db.GetContext(ctx, &service, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &service, err
}

func (r *serviceRepo) GetByCategory(ctx context.Context, categoryID int64) ([]models.Service, error) {
	var services []models.Service
	query := `
		SELECT * FROM services 
		WHERE category_id = $1 
		ORDER BY name
	`
	err := r.db.SelectContext(ctx, &services, query, categoryID)
	return services, err
}

func (r *serviceRepo) GetAll(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	query := `
		SELECT s.*, c.name as category_name 
		FROM services s
		JOIN categories c ON s.category_id = c.category_id
		ORDER BY s.name
	`
	err := r.db.SelectContext(ctx, &services, query)
	return services, err
}

func (r *serviceRepo) Update(ctx context.Context, service *models.Service) error {
	query := `
		UPDATE services 
		SET category_id = :category_id,
		    name = :name,
		    description = :description,
		    is_active = :is_active
		WHERE service_id = :service_id
	`
	result, err := r.db.NamedExecContext(ctx, query, service)
	if err != nil {
		return err
	}
	
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *serviceRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM services WHERE service_id = $1`
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