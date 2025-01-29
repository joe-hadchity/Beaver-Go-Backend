package services

import (
	"server/internal/models"
	"server/internal/repositories"
	"context"
	"errors"
	"fmt"
)

type ServiceService struct {
	serviceRepo  repositories.ServiceRepo
	categoryRepo repositories.CategoryRepo
}

func NewServiceService(
	serviceRepo repositories.ServiceRepo,
	categoryRepo repositories.CategoryRepo,
) *ServiceService {
	return &ServiceService{
		serviceRepo:  serviceRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *ServiceService) CreateService(ctx context.Context, req *models.Service) (*models.Service, error) {
	if req.CategoryID == 0 {
		return nil, errors.New("category ID is required")
	}
	
	if _, err := s.categoryRepo.GetByID(ctx, req.CategoryID); err != nil {
		return nil, fmt.Errorf("invalid category: %w", err)
	}

	if req.Name == "" {
		return nil, errors.New("service name cannot be empty")
	}

	if err := s.serviceRepo.Create(ctx, req); err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}

	return req, nil
}

func (s *ServiceService) GetService(ctx context.Context, id int64) (*models.Service, error) {
	service, err := s.serviceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get service: %w", err)
	}
	return service, nil
}

func (s *ServiceService) ListServices(ctx context.Context) ([]models.Service, error) {
	services, err := s.serviceRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list services: %w", err)
	}
	return services, nil
}

func (s *ServiceService) ListServicesByCategory(ctx context.Context, categoryID int64) ([]models.Service, error) {
	services, err := s.serviceRepo.GetByCategory(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to list services by category: %w", err)
	}
	return services, nil
}

func (s *ServiceService) UpdateService(ctx context.Context, req *models.Service) error {
	if req.ID == 0 {
		return errors.New("invalid service ID")
	}
	
	if err := s.serviceRepo.Update(ctx, req); err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}
	return nil
}

func (s *ServiceService) DeleteService(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("invalid service ID")
	}
	
	if err := s.serviceRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}
	return nil
}