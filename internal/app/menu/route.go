package menu

import (
	m "github.com/alterra-kelompok-1/menu-service/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get, m.NewAuthMiddleware(m.StaffAndAdmin, false).Authenticate)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create)
	g.PUT("/:id", h.Update, m.NewAuthMiddleware(m.StaffAndAdmin, false).Authenticate)
	g.DELETE("/:id", h.Delete, m.NewAuthMiddleware(m.StaffAndAdmin, false).Authenticate)

}
