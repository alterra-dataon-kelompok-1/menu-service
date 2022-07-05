package dto

type CreateMenuRequest struct {
	Name        string `json:"name" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	Description string `json:"description"`
}

type UpdateMenuRequest struct {
	Name        *string `json:"name"`
	Stock       *int    `json:"stock"`
	Description *string `json:"description"`
}
