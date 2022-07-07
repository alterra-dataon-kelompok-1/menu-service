package dto

type CreateMenuRequest struct {
	Name        string `json:"name" validate:"required"`
	Stock       int64  `json:"stock" validate:"required"`
	Description string `json:"description"`
}

type UpdateMenuRequest struct {
	Name        *string `json:"name"`
	Stock       *int64  `json:"stock"`
	Description *string `json:"description"`
}
