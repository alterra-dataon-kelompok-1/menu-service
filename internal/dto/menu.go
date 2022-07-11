package dto

import "github.com/google/uuid"

type CreateMenuRequest struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;uuid_generate_v4;primaryKey"`
	MenuCategoryID int       `json:"menu_category_id" gorm:"not null" validation:"required"`
	Name           string    `json:"name" validation:"required"`
	Description    string    `json:"description" validation:"required"`
	ImageUrl       string    `json:"image_url" gorm:"not null"`
	Price          float64   `json:"price" gorm:"not null" validation:"required"`
	InStock        int64     `json:"in_stock" validation:"required"`
}

type UpdateMenuRequest struct {
	ID             uuid.UUID `json:"id" gorm:"primaryKey"`
	MenuCategoryID *int      `json:"menu_category_id"`
	Name           *string   `json:"name"`
	Description    *string   `json:"description"`
	ImageUrl       *string   `json:"image_url"`
	InStock        *int64    `json:"in_stock"`
}
