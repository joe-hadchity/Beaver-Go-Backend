package models

type Service struct {
    ID          int64  `json:"id" db:"service_id"`
    CategoryID  int64  `json:"category_id" db:"category_id"`
    Name        string `json:"name" db:"name"`
    Description string `json:"description" db:"description"`
    IsActive    bool   `json:"is_active" db:"is_active"`
}