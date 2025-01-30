package models

type Category struct {
        ID          int64  `json:"id" db:"category_id"`
        Name        string `json:"name" db:"name"`
        Description string `json:"description" db:"description"`
        IsActive    bool   `json:"is_active" db:"is_active"`
    }