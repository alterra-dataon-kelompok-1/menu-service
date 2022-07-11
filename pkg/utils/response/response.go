package response

import "github.com/alterra-kelompok-1/menu-service/internal/dto"

type Meta struct {
	Success bool                `json:"success" default:"true"`
	Message string              `json:"message" default:"true"`
	Info    *dto.PaginationInfo `json:"info"`
}
